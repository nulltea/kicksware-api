using System.Runtime.Serialization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace Core.Entities.Users
{
	public class Settings
	{
		[JsonConverter(typeof(StringEnumConverter))]
		public Theme Theme { get; set; } = Theme.Dark;
	}

	[JsonConverter(typeof(StringEnumConverter))]
	public enum Theme
	{
		[EnumMember(Value = "dark")]
		Dark,

		[EnumMember(Value = "light")]
		Light
	}
}