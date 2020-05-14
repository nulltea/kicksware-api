using System;
using Core.Entities.Users;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.DataProtection;
using Microsoft.Extensions.DependencyInjection;
using Web.Handlers.Authentication;

namespace Web.Container.Factory
{
	public partial class ServiceFactory
	{
		public static SecureDataFormat<AuthToken> ProvideSecureTokenFormat(IServiceProvider serviceProvider)
		{
			var protectionProvider = serviceProvider.GetService<IDataProtectionProvider>();
			var dataProtector = protectionProvider.CreateProtector("Web.Auth.MiddlewareAuthHandler", MiddlewareAuthDefaults.AuthenticationScheme, "v2");
			return new SecureDataFormat<AuthToken>(new AuthTokenSerializer(), dataProtector);
		}
	}
}