using Microsoft.AspNetCore.Http;

namespace Web.Handlers.Authentication
{
	public static class MiddlewareAuthDefaults
	{
		public const string AuthenticationScheme = "Middleware.Identity";

		public const string DisplayName = "Authentication Middleware";

		public static readonly PathString LoginPath = "/Account/Login";

		public static readonly PathString LogoutPath = "/Account/Logout";

		public static readonly PathString UnauthorisedPath = "/Account/Unauthorised";

		public const string ReturnUrlParameter = "ReturnUrl";

		public const string CookiePrefix = "Kicksware.";

		public const string TokenPrefix = "auth.token";

		public const string RemoteAuthSchemeKey = ".AuthScheme";

		public static string AuthCookieName => string.Concat(CookiePrefix, AuthenticationScheme);
	}
}