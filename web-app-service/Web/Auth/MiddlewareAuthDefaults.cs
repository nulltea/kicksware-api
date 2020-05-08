using Microsoft.AspNetCore.Http;

namespace Web.Auth
{
	public static class MiddlewareAuthDefaults
	{
		public const string AuthenticationScheme = "MIDDLEWARE";

		public const string DisplayName = "Authentication Middleware";

		public static readonly PathString LoginPath = "/Account/Login";

		public static readonly PathString LogoutPath = "/Account/Logout";

		public static readonly PathString UnauthorisedPath = "/Account/Unauthorised";

		public static readonly string ReturnUrlParameter = "ReturnUrl";

		public static readonly string CookiePrefix = ".AspNetCore.";
	}
}