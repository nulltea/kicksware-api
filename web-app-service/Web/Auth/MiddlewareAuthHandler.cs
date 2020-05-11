using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Net;
using System.Security.Claims;
using System.Text.Encodings.Web;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Http.Features;
using Microsoft.AspNetCore.Identity;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using Microsoft.IdentityModel.Tokens;
using Web.Utils.Extensions;

namespace Web.Auth
{
	public class MiddlewareAuthHandler : SignInAuthenticationHandler<MiddlewareAuthOptions>
	{
		private Task<AuthenticateResult> _getTokenTask;

		private readonly IAuthService _service;

		private readonly UserManager<User> _userManager;

		private const string TokenPropertyName = "AUTHTOKEN";

		public MiddlewareAuthHandler(UserManager<User> userManager, IOptionsMonitor<MiddlewareAuthOptions> options, IAuthService service,
									ILoggerFactory logger, UrlEncoder encoder, ISystemClock clock)
			: base(options, logger, encoder, clock)
		{
			_userManager = userManager;
			_service = service;
		}

		protected override async Task<AuthenticateResult> HandleAuthenticateAsync()
		{
			if (_getTokenTask != null) return await _getTokenTask;

			var cookieToken = ReadCookieToken(out var token);

			if (cookieToken.Succeeded)
			{
				_getTokenTask = Task.FromResult(cookieToken);
				return cookieToken;
			}

			// if (!await _service.ValidateTokenAsync(token)) return AuthenticateResult.Fail("Token not valid");

			var fetchedToken = await RequestGuestToken();

			if (!fetchedToken.Succeeded) return fetchedToken;

			StoreCookieToken(fetchedToken.Ticket.RetrieveToken());

			_getTokenTask = Task.FromResult(fetchedToken);
			return fetchedToken;
		}

		protected override async Task HandleChallengeAsync(AuthenticationProperties properties)
		{
			var authResult = await HandleAuthenticateOnceSafeAsync();

			if (!authResult.Succeeded) Response.StatusCode = (int) HttpStatusCode.Unauthorized;
		}

		protected override async Task HandleSignInAsync(ClaimsPrincipal claims, AuthenticationProperties properties)
		{
			if (claims == null) throw new ArgumentNullException(nameof(claims));

			var user = await ProvideClaimsUser(claims);

			var fetchedToken = await RequestLogin(user);

			if (!fetchedToken.Succeeded) return;

			StoreCookieToken(fetchedToken.Ticket.RetrieveToken());
		}

		protected override Task HandleSignOutAsync(AuthenticationProperties properties)
		{
			var result = ReadCookieToken(out var token);

			if (!result.Succeeded) return Task.CompletedTask;

			_service.LogoutAsync(token);

			RemoveCookieToken();
			return Task.CompletedTask;
		}

		private AuthenticateResult ReadCookieToken(out AuthToken token)
		{
			var cookie = Options.CookieManager.GetRequestCookie(Context, Options.Cookie.Name);

			token = Options.TokenDataFormat.Unprotect(cookie, Context.GetTlsTokenBinding());

			if (token is null) return AuthenticateResult.Fail("Unprotect token failed");

			if (string.IsNullOrEmpty(token)) return AuthenticateResult.NoResult();

			return AuthenticateResult.Success(ProvideTokenAuthTicket(token));
		}

		private async Task<AuthenticateResult> RequestGuestToken()
		{
			if (_service is null) return AuthenticateResult.Fail("Auth service not configured");

			var token = await _service.GuestAsync();

			return string.IsNullOrEmpty(token)
				? AuthenticateResult.Fail("Access forbidden")
				: AuthenticateResult.Success(ProvideTokenAuthTicket(token));
		}

		private async Task<AuthenticateResult> RequestLogin(AuthCredentials credentials)
		{
			if (_service is null) return AuthenticateResult.Fail("Auth service not configured");

			var token = await _service.LoginAsync(credentials);

			return string.IsNullOrEmpty(token)
				? AuthenticateResult.Fail("Access forbidden")
				: AuthenticateResult.Success(ProvideTokenAuthTicket(token));
		}

		private void StoreCookieToken(AuthToken token)
		{
			var cookieValue = Options.TokenDataFormat.Protect(token, Context.GetTlsTokenBinding());

			Options.CookieManager.AppendResponseCookie(
				Context,
				Options.Cookie.Name,
				cookieValue,
				Options.Cookie.Build(Context));
		}

		private void RemoveCookieToken()
		{
			Options.CookieManager.DeleteCookie(
				Context,
				Options.Cookie.Name,
				Options.Cookie.Build(Context));
		}

		private AuthenticationTicket ProvideTokenAuthTicket(AuthToken token)
		{
			var claims = ReadTokenPayload(token);
			if (claims is null) throw new NullReferenceException(nameof(claims));

			var ticket = new AuthenticationTicket(claims, token.ToAuthProperties(), Scheme.Name);
			return ticket;
		}

		private Task<User> ProvideClaimsUser(ClaimsPrincipal claims) => _userManager.GetUserAsync(claims);

		private ClaimsPrincipal ReadTokenPayload(AuthToken token)
		{
			var jwtHandler = new JwtSecurityTokenHandler {MapInboundClaims = true};
			var claims = jwtHandler.ReadJwtToken(token).Claims?.ToList();

			if (claims is null || !claims.Any()) return null;

			for (var i = 0; i < claims.Count; i++)
			{
				var claim = claims[i];
				if (!jwtHandler.InboundClaimTypeMap.TryGetValue(claim.Type, out var mappedType)) continue;
				claims[i] = new Claim(mappedType, claim.Value, claim.Issuer, claim.Issuer);
			}

			var identity = new ClaimsIdentity(claims, MiddlewareAuthDefaults.AuthenticationScheme);

			return new ClaimsPrincipal(identity);
		}

		private void CheckForRefresh(AuthenticationTicket token)
		{
			var currentUtc = Clock.UtcNow;
			var expiresUtc = token.Properties.ExpiresUtc;
			var allowRefresh =  token.Properties.AllowRefresh ?? true;
			if (!allowRefresh) return;
			if (expiresUtc < currentUtc)
			{
				//_service.RefreshToken(ref token);
				StoreCookieToken(token.RetrieveToken());
			}
		}
	}
}