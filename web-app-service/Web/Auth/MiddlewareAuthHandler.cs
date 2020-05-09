using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Claims;
using System.Text.Encodings.Web;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Http.Features;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;

namespace Web.Auth
{
	public class MiddlewareAuthHandler : SignInAuthenticationHandler<MiddlewareAuthOptions>
	{
		private Task<AuthenticateResult> _getTokenTask;

		private readonly IAuthService _service;

		private const string TokenPropertyName = "AUTHTOKEN";

		public MiddlewareAuthHandler(IOptionsMonitor<MiddlewareAuthOptions> options, IAuthService service,
									ILoggerFactory logger, UrlEncoder encoder, ISystemClock clock)
			: base(options, logger, encoder, clock)
		{
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

			var fetchedToken = await FetchGuestToken();

			if (!fetchedToken.Succeeded) return fetchedToken;

			StoreCookieToken(RetrieveFromTicket(fetchedToken.Ticket));

			_getTokenTask = Task.FromResult(fetchedToken);
			return fetchedToken;
		}

		protected override async Task HandleChallengeAsync(AuthenticationProperties properties)
		{
			throw new NotImplementedException($"{nameof(HandleChallengeAsync)} not implemented");
		}

		protected override async Task HandleSignInAsync(ClaimsPrincipal user, AuthenticationProperties properties)
		{
			if (user == null) throw new ArgumentNullException(nameof(user));
			// properties = properties ?? new AuthenticationProperties();
			var credentials = ProvideClaimsCredentials(user);

			var fetchedToken = await FetchLoginToken(credentials);

			if (!fetchedToken.Succeeded) return;

			StoreCookieToken(RetrieveFromTicket(fetchedToken.Ticket));
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

			token = Options.TokenDataFormat.Unprotect(cookie, GetTlsTokenBinding());

			if (token is null) return AuthenticateResult.Fail("Unprotect token failed");

			if (string.IsNullOrEmpty(token)) return AuthenticateResult.NoResult();

			return AuthenticateResult.Success(ProvideTokenAuthTicket(token));
		}

		private async Task<AuthenticateResult> FetchGuestToken()
		{
			if (_service is null) return AuthenticateResult.Fail("Auth service not configured");

			var token = await _service.GuestAsync();

			return string.IsNullOrEmpty(token)
				? AuthenticateResult.Fail("Access forbidden")
				: AuthenticateResult.Success(ProvideTokenAuthTicket(token));
		}

		private async Task<AuthenticateResult> FetchLoginToken(AuthCredentials credentials)
		{
			if (_service is null) return AuthenticateResult.Fail("Auth service not configured");

			var token = await _service.LoginAsync(credentials);

			return string.IsNullOrEmpty(token)
				? AuthenticateResult.Fail("Access forbidden")
				: AuthenticateResult.Success(ProvideTokenAuthTicket(token));
		}

		private void StoreCookieToken(AuthToken token)
		{
			var cookieValue = Options.TokenDataFormat.Protect(token, GetTlsTokenBinding());

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
			var ticket = new AuthenticationTicket(ClaimsPrincipal.Current, ProvideTokenAuthProperties(token),
				Scheme.Name);
			return ticket;
		}

		private AuthenticationProperties ProvideTokenAuthProperties(AuthToken token) =>
			new AuthenticationProperties(new Dictionary<string, string>
			{
				{TokenPropertyName, token.Token}
			});

		private AuthToken RetrieveFromTicket(AuthenticationTicket ticket) => ticket.Properties.GetParameter<AuthToken>(TokenPropertyName);

		private AuthCredentials ProvideClaimsCredentials(ClaimsPrincipal claims)
		{

			var claimsMap = claims.Claims.ToDictionary(claim => claim.Type, claim => claim.Value);
			return new AuthCredentials(claimsMap[ClaimTypes.Email], claimsMap["pwd"]);
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
				StoreCookieToken(RetrieveFromTicket(token));
			}
		}

		private string GetTlsTokenBinding()
		{
			var binding = Context.Features.Get<ITlsTokenBindingFeature>()?.GetProvidedTokenBindingId();
			return binding == null ? null : Convert.ToBase64String(binding);
		}
	}
}