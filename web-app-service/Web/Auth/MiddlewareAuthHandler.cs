using System;
using System.Text.Encodings.Web;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authentication;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;

namespace Web.Auth
{
	public class MiddlewareAuthHandler : AuthenticationHandler<MiddlewareAuthOptions>
	{
		private Task<AuthenticateResult> _readCookieTask;

		private IAuthService _service;

		public MiddlewareAuthHandler(IOptionsMonitor<MiddlewareAuthOptions> options, IAuthService service,
									ILoggerFactory logger, UrlEncoder encoder, ISystemClock clock)
			: base(options, logger, encoder, clock)
		{
			_service = service;
		}

		protected override async Task<AuthenticateResult> HandleAuthenticateAsync()
		{
			var result = await EnsureCookieToken();
			return AuthenticateResult.Fail("No implemented");
		}

		protected override async Task HandleChallengeAsync(AuthenticationProperties properties)
		{
		}

		private Task<AuthenticateResult> EnsureCookieToken()
		{
			// We only need to read the tocken once
			if (_readCookieTask == null)
			{
				_readCookieTask = ReadCookieToken();
			}
			return _readCookieTask;
		}

		private async Task<AuthenticateResult> ReadCookieToken()
		{
			throw new NotImplementedException("ReadCookieTicket not implemented");
		}

		private void CheckForRefresh(AuthToken token)
		{
			var currentUtc = Clock.UtcNow;
			var expiresUtc = token.ExpiresUtc;
			var allowRefresh = token.AllowRefresh;
			if (!allowRefresh) return;
			if (expiresUtc < currentUtc)
			{
				_service.RefreshToken(ref token);
				// TODO cookie token
			}
		}
	}
}