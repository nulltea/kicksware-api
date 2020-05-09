using Microsoft.AspNetCore.Mvc;
using Web.Controllers;

namespace Web.Utils.UrlHelpers
{
	public static class UrlHelperExtensions
	{
		public static string EmailConfirmationLink(this IUrlHelper urlHelper, string userId, string code, string scheme)
		{
			return urlHelper.Action(action: nameof(AuthController.ConfirmEmail), "Auth",
				new {userId, code}, scheme);
		}

		public static string ResetPasswordCallbackLink(this IUrlHelper urlHelper, string userId, string code,
														string scheme)
		{
			return urlHelper.Action(action: nameof(AuthController.ResetPassword), "Auth",
				new {userId, code}, scheme);
		}
	}
}