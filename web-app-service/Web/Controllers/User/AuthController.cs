using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.Serialization;
using System.Security.Claims;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Services;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Web.Models.Auth;
using Web.Utils.Extensions;
using Web.Utils.UrlHelpers;

namespace Web.Controllers
{
	[Authorize(Policy = "NotGuest")]
	public class AuthController : Controller
	{
		private readonly IUserService _service;

		private readonly UserManager<User> _userManager;

		private readonly SignInManager<User> _signInManager;

		private readonly ILogger _logger;

		public AuthController(IUserService service, UserManager<User> userManager, SignInManager<User> signInManager, ILogger<AuthController> logger)
		{
			_service = service;
			_userManager = userManager;
			_signInManager = signInManager;
			_logger = logger;
		}

		[HttpGet]
		public async Task<IActionResult> Auth()
		{
			var user = await _userManager.GetUserAsync(HttpContext.User);
			if (!_signInManager.IsSignedIn(HttpContext.User) ||
				user is null || !user.Confirmed)
			{
				var content = await this.RenderViewAsync("_AuthDialogPartial",
					new AuthCommonViewModel
					{
						Email = user?.Email,
						UserName = user?.Username,
						AwaitVerification = !(user?.Confirmed ?? true)
					}, true);
				return Json(new
				{
					IsLogedIn = false,
					Content = content
				});
			}

			return Json(new
			{
				IsLogedIn = true,
				RedirectUrl = Url.Action("Profile", "Profile")
			});
		}

		[HttpPost]
		public Task<IActionResult> Auth(AuthCommonViewModel model, AuthMode mode)
		{
			if (mode == AuthMode.Login) return Login(model);

			return SignUp(model);
		}

		[HttpGet]
		[AllowAnonymous]
		public async Task<IActionResult> Login(string returnUrl = default)
		{
			await HttpContext.SignOutAsync(IdentityConstants.ExternalScheme);

			ViewData["ReturnUrl"] = returnUrl;
			return View();
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> Login(LoginViewModel model, string returnUrl = default)
		{
			ViewData["ReturnUrl"] = returnUrl;
			if (ModelState.IsValid)
			{
				var result = await _signInManager.PasswordSignInAsync(model.Email, model.Password, model.RememberMe,
					lockoutOnFailure: false);
				if (result.Succeeded)
				{
					_logger.LogInformation("User logged in.");
					return RedirectToLocal(returnUrl);
				}

				ModelState.AddModelError(string.Empty, "Invalid login attempt.");
				return View(model);
			}

			return View(model);
		}


		[HttpGet]
		[AllowAnonymous]
		public async Task<IActionResult> LoginWithRecoveryCode(string returnUrl = default)
		{
			var user = await _signInManager.GetTwoFactorAuthenticationUserAsync();
			if (user is null) throw new ApplicationException("Unable to load two-factor authentication user.");
			ViewData["ReturnUrl"] = returnUrl;
			return View();
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> LoginWithRecoveryCode(RecoveryCodeViewModel model, string returnUrl = default)
		{
			if (!ModelState.IsValid) return View(model);

			var user = await _signInManager.GetTwoFactorAuthenticationUserAsync();

			if (user == null) throw new ApplicationException("Unable to load two-factor authentication user.");

			var recoveryCode = model.RecoveryCode.Replace(" ", string.Empty);

			var result = await _signInManager.TwoFactorRecoveryCodeSignInAsync(recoveryCode);

			if (result.Succeeded) return RedirectToLocal(returnUrl);

			ModelState.AddModelError(string.Empty, "Invalid recovery code entered.");

			return View();
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ContinueSignUp(SignUpViewModel model, string returnUrl = default)
		{
			ViewData["ReturnUrl"] = returnUrl;
			return View("SignUp", model);
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> SignUp(SignUpViewModel model, string returnUrl = default)
		{
			ViewData["ReturnUrl"] = returnUrl;

			// if (!ModelState.IsValid) return View(model);

			var user = new User { Email = model.Email};
			var result = await _userManager.CreateAsync(user, model.Password);

			if (result.Succeeded)
			{
				_logger.LogInformation("User created a new account with password.");

				user = await _userManager.FindByEmailAsync(model.Email);
				await _signInManager.SignInWithClaimsAsync(user, false, new []
				{
					new Claim(ClaimTypes.Email, user.Email),
					new Claim(ClaimTypes.Hash, user.PasswordHash)
				});
				_logger.LogInformation("User created a new account with password.");
				return Json(new {success = true});
			}

			AddErrors(result);
			return Json(new {success = false});
		}


		public IActionResult Facebook()
		{
			var authProperties = new AuthenticationProperties {RedirectUri = Url.Action("Index", "Home")};
			return Challenge(authProperties, "Facebook");
		}

		public IActionResult Google()
		{
			var authProperties = new AuthenticationProperties {RedirectUri = Url.Action("Index", "Home")};
			return Challenge(authProperties, "Google");
		}

		[HttpPost]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> Logout()
		{
			await _signInManager.SignOutAsync();
			_logger.LogInformation("User logged out.");
			return RedirectToAction(nameof(HomeController.Index), "Home");
		}

		[HttpGet]
		[AllowAnonymous]
		public async Task<IActionResult> ResendEmail(string userId, string code)
		{
			if (userId is null || code is null) return RedirectToAction(nameof(HomeController.Index), "Home");

			var user = await _userManager.FindByIdAsync(userId);
			if (user == null) throw new ApplicationException($"Unable to load user with ID '{userId}'.");

			var result = await _userManager.ConfirmEmailAsync(user, code);
			return View(result.Succeeded ? "ConfirmEmail" : "Error");
		}

		[HttpGet]
		[AllowAnonymous]
		public async Task<IActionResult> ConfirmEmail(string userId, string code)
		{
			if (userId is null || code is null) return RedirectToAction(nameof(HomeController.Index), "Home");

			var user = await _userManager.FindByIdAsync(userId);
			if (user == null) throw new ApplicationException($"Unable to load user with ID '{userId}'.");

			var result = await _userManager.ConfirmEmailAsync(user, code);
			return View(result.Succeeded ? "ConfirmEmail" : "Error");
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ForgotPassword()
		{
			return View();
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> ForgotPassword(ForgotPasswordViewModel model)
		{
			if (!ModelState.IsValid) return View(model);

			var user = await _userManager.FindByEmailAsync(model.Email);
			if (user is null || !await _userManager.IsEmailConfirmedAsync(user))
			{
				return RedirectToAction(nameof(ForgotPasswordConfirmation));
			}

			var code = await _userManager.GeneratePasswordResetTokenAsync(user);
			var callbackUrl = Url.ResetPasswordCallbackLink(user.UniqueID, code, Request.Scheme);
			await _service.SendResetPasswordEmailAsync(user.UniqueID, callbackUrl);
			return RedirectToAction(nameof(ForgotPasswordConfirmation));
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ForgotPasswordConfirmation()
		{
			return View();
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ResetPassword(string code = null)
		{
			if (code is null) throw new ApplicationException("A code must be supplied for password reset.");
			var model = new ResetPasswordViewModel {Code = code};
			return View(model);
		}

		[HttpPost]
		[AllowAnonymous]
		[ValidateAntiForgeryToken]
		public async Task<IActionResult> ResetPassword(ResetPasswordViewModel model)
		{
			if (!ModelState.IsValid) return View(model);

			var user = await _userManager.FindByEmailAsync(model.Email);
			if (user is null) return RedirectToAction(nameof(ResetPasswordConfirmation));

			var result = await _userManager.ResetPasswordAsync(user, model.Code, model.Password);
			if (result.Succeeded) return RedirectToAction(nameof(ResetPasswordConfirmation));

			AddErrors(result);
			return View();
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ResetPasswordConfirmation() => View();


		[HttpGet]
		public IActionResult AccessDenied(string fromAction)
		{
			TempData.Add("locked", true);
			return View();
		}

		#region Helpers


		private void AddErrors(IdentityResult result)
		{
			foreach (var error in result.Errors)
			{
				ModelState.AddModelError(string.Empty, error.Description);
			}
		}

		private IActionResult RedirectToLocal(string returnUrl)
		{
			if (Url.IsLocalUrl(returnUrl))
			{
				return Redirect(returnUrl);
			}
			else
			{
				return RedirectToAction(nameof(HomeController.Index), "Home");
			}
		}

		#endregion

		public enum AuthMode
		{
			SingUp,

			Login
		}
	}
}