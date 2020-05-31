using System;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Facebook;
using Microsoft.AspNetCore.Authentication.Google;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Web.Models.Auth;
using Web.Utils.Extensions;
using Web.Utils.UrlHelpers;

namespace Web.Controllers
{
	public class AuthController : Controller
	{
		private readonly IUserService _service;

		private readonly UserManager<User> _userManager;

		private readonly SignInManager<User> _signInManager;

		private readonly ILogger _logger;

		public AuthController(IUserService service, UserManager<User> userManager, SignInManager<User> signInManager,
							ILogger<AuthController> logger)
		{
			_service = service;
			_userManager = userManager;
			_signInManager = signInManager;
			_logger = logger;
		}

		#region Main auth operations

		[HttpGet]
		public async Task<IActionResult> Auth(User user = default)
		{
			if (string.IsNullOrEmpty(user.UniqueID)) user = null;
			user ??= await _userManager.GetUserAsync(HttpContext.User);
			if (!_signInManager.IsSignedIn(HttpContext.User) || user is null || !user.Confirmed)
			{
				var verifyPending = !(user?.Confirmed ?? true);

				var content = await this.RenderViewAsync("_AuthDialogPartial",
					new AuthCommonViewModel
					{
						Email = user?.Email, UserName = user?.Username, VerifyPending = verifyPending
					}, true);

				return Json(new
				{
					Success = true,
					Logged = false,
					Content = content,
					VerifyPending = verifyPending,
					RedirectURL = Url.Action("Index", "Home")
				});
			}

			return Json(new {Success = true, Logged = true, RedirectUrl = Url.Action("Profile", "Profile")});
		}

		[HttpPost]
		public Task<IActionResult> Auth(AuthCommonViewModel model, AuthMode mode)
		{
			return mode == AuthMode.Login ? Login(model) : SignUp(model);
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> SignUp(SignUpViewModel model, string returnUrl = default)
		{
			ViewData["ReturnUrl"] = returnUrl;

			// TODO if (!ModelState.IsValid) return View(model);

			var user = new User {Email = model.Email};
			var result = await _userManager.CreateAsync(user, model.Password);

			if (!result.Succeeded) return Json(new
			{
				Success = false,
				Error = result.Errors.FirstOrDefault()?.Description
			});

			user = await _userManager.FindByEmailAsync(model.Email);
			if (user is null) return Json(new
			{
				Success = false,
				Error = "Something went wrong during creation your account. Please try again soon"
			});

			_logger.LogInformation($"User {user.Username} created a new account with password");
			await _signInManager.SignInWithClaimsAsync(user, false, user.ExtractCredentials());
			_logger.LogInformation($"User {user.Username} signed up with new account");

			if (!await SendEmailConfirmationAsync(user)) return Json(new
			{
				Success = false,
				Error = "Something went wrong during sending you an email. Please try again soon"
			});

			return await Auth(user);
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> Login(LoginViewModel model, string returnUrl = default)
		{
			ViewData["ReturnUrl"] = returnUrl;
			var user = await _userManager.FindByEmailAsync(model.Email);
			if (user is null)
				return Json(new
				{
					Success = false,
					Error = $"User with email {model.Email} was not found.\nPlease check your credentials"
				});

			if (!await _userManager.CheckPasswordAsync(user, model.Password))
			{
				return Json(new {Success = false, Error = "The password you entered is incorrect. Please try again"});
			}

			var result = await _signInManager.PasswordSignInAsync(user, model.Password, model.RememberMe, false);
			if (result.Succeeded)
			{
				_logger.LogInformation($"User {user.Username} logged");
				return await Auth(user);
			}

			return Json(new
			{
				Success = false,
				Error = "Something went wrong during logging in. Please try again soon"
			});
		}

		public IActionResult Facebook()
		{
			var authProperties = new AuthenticationProperties {RedirectUri = Url.Action("Profile", "Profile")};
			return Challenge(authProperties, FacebookDefaults.AuthenticationScheme);
		}

		public IActionResult Google()
		{
			var authProperties = new AuthenticationProperties {RedirectUri = Url.Action("Profile", "Profile")};
			return Challenge(authProperties, GoogleDefaults.AuthenticationScheme);
		}

		public async Task<IActionResult> Logout()
		{
			await _signInManager.SignOutAsync();
			_logger.LogInformation("User logged out.");
			return RedirectToAction("Index", "Home");
		}

		#endregion

		#region Additional auth operations

		[HttpGet]
		[AllowAnonymous]
		public async Task<IActionResult> ResendEmail(string email)
		{
			var user = await _userManager.FindByEmailAsync(email);
			user ??= await _userManager.GetUserAsync(HttpContext.User);

			if (user is null) return Json(new
			{
				Success = false,
				Error = "Cannot find specified user. Please check your email and try again"
			});

			if (!await SendEmailConfirmationAsync(user)) return Json(new
			{
				Success = false,
				Error = "Something went wrong during sending you an email. Please try again soon"
			});

			return Json(new {Success = true});
		}

		[HttpGet]
		[AllowAnonymous]
		public async Task<IActionResult> ConfirmEmail(string userId, string code)
		{
			if (userId is null || code is null) return RedirectToAction(nameof(HomeController.Index), "Home");

			var user = await _userManager.FindByIdAsync(userId);
			if (user == null) throw new ApplicationException($"Unable to load user with ID '{userId}'.");

			var result = await _userManager.ConfirmEmailAsync(user, code);
			return View(result.Succeeded ? "ConfirmEmail" : "Error", user);
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> ForgotPassword(LoginViewModel model)
		{
			var user = await _userManager.FindByEmailAsync(model.Email);
			if (user is null || !await _userManager.IsEmailConfirmedAsync(user))
			{
				return Json(new
				{
					Success = false,
					Error = $"User with email {model.Email} was not found.\nPlease check your credentials"
				});
			}

			var code = await _userManager.GeneratePasswordResetTokenAsync(user);
			var callbackUrl = Url.ResetPasswordCallbackLink(user.UniqueID, code, Request.Scheme);
			if (!await _service.SendResetPasswordEmailAsync(user.UniqueID, callbackUrl))
			{
				return Json(new
				{
					Success = false,
					Error = "Something went wrong during sending you an email. Please try again soon"
				});
			}

			return Json(new
			{
				Success = true,
				Content = await this.RenderViewAsync("_ForgotPassword",
					user, true)
			});
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ResetPassword(string code = null)
		{
			if (code is null) throw new ApplicationException("A code must be supplied for password reset");
			var model = new ResetPasswordViewModel {Code = code};
			return View(model);
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> ResetPassword(ResetPasswordViewModel model)
		{
			var user = await _userManager.FindByEmailAsync(model.Email);
			if (user is null) return Json(new
			{
				Success = false,
				Error = $"User with email {model.Email} was not found.\nPlease check your credentials"
			});

			if (string.IsNullOrWhiteSpace(model.Password)) return Json(new
			{
				Success = false,
				Error = "Would you like me to choose your new password? Spoiler: I can't do that..."
			});

			if (!model.Password.Equals(model.ConfirmPassword))
			{
				return Json(new
				{
					Success = false,
					Error = "Password confirmation and Password must match"
				});
			}

			var result = await _userManager.ResetPasswordAsync(user, model.Code, model.Password);
			if (!result.Succeeded) return Json(new
			{
				Success = false,
				Error = result.Errors.FirstOrDefault()?.Description
			});

			return Json(new
			{
				Success = true,
				RedirectURL = Url.Action("Index", "Home"),
				Content = await this.RenderViewAsync("_ResetPasswordConfirm", new AuthCommonViewModel
				{
					Email = model.Email,
					AuthSign = true,
				}, true)
			});
		}

		[HttpGet]
		public IActionResult AccessDenied()
		{
			return View();
		}

		#endregion

		#region Service & Helpers

		private async Task<bool> SendEmailConfirmationAsync(User user)
		{
			var code = await _userManager.GenerateEmailConfirmationTokenAsync(user);
			var callbackUrl = Url.EmailConfirmationLink(user.UniqueID, code, Request.Scheme);
			if (!await _service.SendEmailConfirmationAsync(user.UniqueID, callbackUrl)) return false;

			_logger.LogInformation($"Email conformation was sent to user {user.Username}");
			return true;
		}


		public enum AuthMode
		{
			SingUp,

			Login
		}

		#endregion
	}
}