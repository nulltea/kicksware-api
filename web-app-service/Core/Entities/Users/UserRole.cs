using System.Runtime.Serialization;

namespace Core.Entities.Users
{
	public enum UserRole
	{
		[EnumMember(Value="")]
		Regular,

		[EnumMember(Value="gst")]
		Guest,

		[EnumMember(Value="adm")]
		Admin,
	}
}