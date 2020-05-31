using System.Runtime.Serialization;
using System.Text.Json.Serialization;
using Newtonsoft.Json.Converters;

namespace Core.Entities.Users
{
	[JsonConverter(typeof(StringEnumConverter))]
	public enum UserProvider
	{
		[EnumMember(Value = "")]
		Internal,

		[EnumMember(Value = "facebook")]
		Facebook,

		[EnumMember(Value = "google")]
		Google,

		[EnumMember(Value = "apple")]
		Apple
	}
}