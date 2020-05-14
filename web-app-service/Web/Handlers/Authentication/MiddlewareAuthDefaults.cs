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

		public static readonly string ReturnUrlParameter = "ReturnUrl";

		public static readonly string CookiePrefix = "Kicksware.";

		public static readonly string TokenPrefix = "auth.token";

		public static string AuthCookieName => string.Concat(CookiePrefix, AuthenticationScheme);
	}
}