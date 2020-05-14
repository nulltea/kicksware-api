using Core.Entities.Users;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Options;

namespace Web.Handlers.Authentication
{
	public class MiddlewareAuthPostConfigureOptions : IPostConfigureOptions<MiddlewareAuthOptions>
	{
		private readonly ISecureDataFormat<AuthToken> _secureTokenFormat;
		public MiddlewareAuthPostConfigureOptions(ISecureDataFormat<AuthToken> secureTokenFormat) => _secureTokenFormat = secureTokenFormat;

		public void PostConfigure(string name, MiddlewareAuthOptions options)
		{
			options.CookieManager ??= new ChunkingCookieManager();

			options.Cookie ??= new RequestPathBaseCookieBuilder
			{
				SameSite = SameSiteMode.Lax,
				HttpOnly = false,
				SecurePolicy = CookieSecurePolicy.SameAsRequest,
				IsEssential = true,
			};

			if (string.IsNullOrEmpty(options.Cookie.Name)) options.Cookie.Name = MiddlewareAuthDefaults.AuthCookieName;

			if (!options.LoginPath.HasValue) options.LoginPath = MiddlewareAuthDefaults.LoginPath;

			if (!options.LogoutPath.HasValue) options.LogoutPath = MiddlewareAuthDefaults.LogoutPath;

			if (!options.UnauthorisedPath.HasValue) options.UnauthorisedPath = MiddlewareAuthDefaults.UnauthorisedPath;

			options.TokenDataFormat ??= _secureTokenFormat;
		}
	}
}