using Core.Attributes;

namespace Core.Entities.Users
{
	[EntityService(Resource = "api/users")]
	public class User : AuthCredentials, IBaseEntity
	{
		public string UniqueID { get; set; }

		public string Username { get; set; }

		public string FirstName { get; set; }

		public string LastName { get; set; }

		public string Avatar { get; set; }

		public string PhoneNumber { get; set; }

		public string Location { get; set; }

		public bool Confirmed { get; set; }

		public string Role { get; set; }

		public PaymentInfo PaymentInfo { get; set; } = new PaymentInfo();

		public bool Equals(User other) => other != null && UniqueID == other.UniqueID;
	}
}