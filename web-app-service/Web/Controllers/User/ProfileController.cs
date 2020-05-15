using System.Linq;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;

namespace Web.Controllers
{
	[Authorize(Policy = "NotGuest")]
	public class ProfileController : Controller
	{
		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/profile-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Profile";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "Profile";

		[ViewData]
		public string HeroLogoPath { get; set; }

		private IUserService _service;

		private UserManager<User> _userManager;

		private SignInManager<User> _authManager;

		public ProfileController(IUserService service, UserManager<User> userManager, SignInManager<User> authManager)
		{
			_service = service;
			_userManager = userManager;
			_authManager = authManager;
		}

		[HttpGet]
		[Authorize]
		[Route("profile/{mode?}")]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public async Task<IActionResult> Profile()
		{
			var user = await _userManager.GetUserAsync(HttpContext.User);

			if (user is null) return RedirectToAction("Auth", "Auth");

			if (!string.IsNullOrEmpty(user.FirstName) || !string.IsNullOrEmpty(user.LastName))
			{
				HeroBreadSubTitle = string.Join(" ", user.FirstName, user.LastName);
			}
			else
			{
				HeroBreadSubTitle = user.Username ?? user.Email;
			}

			if (!string.IsNullOrEmpty(user.Avatar)) HeroLogoPath = user.Avatar;

			return View(user);
		}

		[HttpPost]
		[Authorize]
		public async Task<IActionResult> Account(UserViewModel user)
		{
			var result = await _userManager.UpdateAsync(user);

			var updateResult = result.Succeeded
				? FormSubmitResult(SubmitResult.Success, "Great! Account information was successfully updated")
				: FormSubmitResult(SubmitResult.Error, result.Errors.Select(err => err.Description).FirstOrDefault());

			if (!result.Succeeded) return updateResult;

			if (!string.IsNullOrWhiteSpace(user.NewPassword))
			{
				if (!user.NewPassword.Equals(user.ConfirmedPassword))
				{
					return FormSubmitResult(SubmitResult.Error, "Password confirmation and Password must match");
				}

				result = await _userManager.ChangePasswordAsync(user, user.CurrentPassword, user.NewPassword);

				return result.Succeeded
					? FormSubmitResult(SubmitResult.Success, "Nice! Got yourself a new secret password")
					: FormSubmitResult(SubmitResult.Error, result.Errors.Select(err => err.Description).FirstOrDefault());
			}

			return updateResult;
		}

		private JsonResult FormSubmitResult(SubmitResult result, string message) => Json(new
		{
			Result = result.ToString().ToLower(),
			Message = message
		});

		private enum SubmitResult
		{
			Success,

			Error,

			Warning
		}


	}
}