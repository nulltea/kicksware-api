using System.Runtime.Serialization;
using System.Text.Json.Serialization;
using Newtonsoft.Json.Converters;

namespace Core.Entities.Users
{
	[JsonConverter(typeof(StringEnumConverter))]
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