using System.Collections.Generic;
using Core.Attributes;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace Core.Entities.Users
{
	[EntityService(Resource = "users")]
	public class User : AuthCredentials, IBaseEntity
	{
		public string UniqueID { get; set; }

		public string Username { get; set; }

		public string UsernameN { get; set; }

		public string EmailN { get; set; }

		public string FirstName { get; set; }

		public string LastName { get; set; }

		public string Avatar { get; set; }

		public string PhoneNumber { get; set; }

		public string Location { get; set; }

		public bool Confirmed { get; set; }

		[JsonConverter(typeof(StringEnumConverter))]
		public UserRole Role { get; set; }

		public PaymentInfo PaymentInfo { get; set; } = new PaymentInfo();

		public string[] Liked { get; set; }

		public Settings Settings { get; set; } = new Settings();

		[JsonConverter(typeof(StringEnumConverter))]
		public UserProvider Provider { get; set; }

		public Dictionary<UserProvider, string> ConnectedProviders { get; set; } = new Dictionary<UserProvider, string>();

		public bool Equals(User other) => other != null && UniqueID == other.UniqueID;

		public bool IsEmpty() => string.IsNullOrEmpty(UniqueID) && string.IsNullOrEmpty(Email) && string.IsNullOrEmpty(Username);
	}
}