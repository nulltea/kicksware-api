using System.ComponentModel.DataAnnotations;

namespace Web.Models.Auth
{
	public class AuthCommonViewModel
	{
		[Required]
		[EmailAddress]
		[Display(Name = "Email")]
		public string Email { get; set; }

		[Required]
		[Display(Name = "Username")]
		public string UserName { get; set; }

		[Required]
		[StringLength(100, ErrorMessage = "The {0} must be at least {2} and at max {1} characters long.", MinimumLength = 6)]
		[DataType(DataType.Password)]
		[Display(Name = "Password")]
		public string Password { get; set; }

		/// <summary>
		/// Login: Remember me | Sign up: Email notification
		/// </summary>
		public bool AuthSign { get; set; }

		public bool VerifyPending { get; set; }

		public static implicit operator LoginViewModel(AuthCommonViewModel model) =>
			new LoginViewModel
			{
				Email = model.Email,
				Username = model.UserName,
				Password = model.Password,
				RememberMe = model.AuthSign
			};

		public static implicit operator SignUpViewModel(AuthCommonViewModel model) =>
			new SignUpViewModel
			{
				Email = model.Email,
				Username = model.UserName,
				Password = model.Password,
				EmailNotification = model.AuthSign
			};
	}
}