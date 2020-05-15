using System;
using System.Linq;
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
					RedirectUrl = Url.Action("Index", "Home")
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

			await SendEmailConfirmationAsync(user);

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

			return Json(new {Success = false, Error = "Something went wrong during logging in. Please try again soon"});
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

			if (result.Succeeded) return Redirect(returnUrl);

			ModelState.AddModelError(string.Empty, "Invalid recovery code entered.");

			return View();
		}

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

			await SendEmailConfirmationAsync(user);

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

			return View();
		}

		[HttpGet]
		[AllowAnonymous]
		public IActionResult ResetPasswordConfirmation() => View();


		[HttpGet]
		public IActionResult AccessDenied(string fromAction)
		{
			return View();
		}

		#endregion

		#region Service & Helpers

		private async Task SendEmailConfirmationAsync(User user)
		{
			var code = await _userManager.GenerateEmailConfirmationTokenAsync(user);
			var callbackUrl = Url.EmailConfirmationLink(user.UniqueID, code, Request.Scheme);
			await _service.SendEmailConfirmationAsync(user.UniqueID, callbackUrl);

			_logger.LogInformation($"Email conformation was sent to user {user.Username}");

		}

		public enum AuthMode
		{
			SingUp,

			Login
		}

		#endregion
	}
}