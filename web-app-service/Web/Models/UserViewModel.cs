using Core.Entities.Users;

namespace Web.Models
{
	public class UserViewModel : User
	{
		public string CurrentPassword { get; set; }

		public string NewPassword { get; set; }

		public string ConfirmedPassword { get; set; }
	}
}