using System.ComponentModel.DataAnnotations;

namespace Web.Models.Auth
{
	public class ResetPasswordViewModel
	{
		[Required]
		[EmailAddress]
		public string Email { get; set; }

		[DataType(DataType.Password)]
		public string Password { get; set; }

		[DataType(DataType.Password)]
		public string ConfirmPassword { get; set; }

		public string Code { get; set; }
	}
}