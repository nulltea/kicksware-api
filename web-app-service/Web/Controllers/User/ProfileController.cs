using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

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
		public async Task<IActionResult> Account(User user)
		{
			var result = await _userManager.UpdateAsync(user);
			if (result.Succeeded) return Json(new {success = true});
			return Json(new {success = false, errors = result.Errors.Select(err => err.Description)});
		}
	}
}