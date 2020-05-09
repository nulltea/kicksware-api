using Core.Entities.Users;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.DataProtection;
using Microsoft.Extensions.Options;
using Newtonsoft.Json;
using Web.Auth.Service;

namespace Web.Auth
{
	public class MiddlewareAuthPostConfigureOptions : IPostConfigureOptions<MiddlewareAuthOptions>
	{
		private readonly IDataProtectionProvider _dataProtection;
		public MiddlewareAuthPostConfigureOptions(IDataProtectionProvider dataProtection) => _dataProtection = dataProtection;

		public void PostConfigure(string name, MiddlewareAuthOptions options)
		{
			if (options.CookieManager is null) options.CookieManager = new ChunkingCookieManager();

			if (string.IsNullOrEmpty(options.Cookie.Name)) options.Cookie.Name = MiddlewareAuthDefaults.CookiePrefix + name;

			if (!options.LoginPath.HasValue) options.LoginPath = MiddlewareAuthDefaults.LoginPath;

			if (!options.LogoutPath.HasValue) options.LogoutPath = MiddlewareAuthDefaults.LogoutPath;

			if (!options.UnauthorisedPath.HasValue) options.UnauthorisedPath = MiddlewareAuthDefaults.UnauthorisedPath;

			options.DataProtectionProvider ??= _dataProtection;

			if (options.TokenDataFormat is null)
			{
				var dataProtector = options.DataProtectionProvider.CreateProtector("Web.Auth.MiddlewareAuthHandler", name, "v2");
				options.TokenDataFormat = new SecureDataFormat<AuthToken>(new AuthTokenSerializer(), dataProtector);
			}
		}
	}
}