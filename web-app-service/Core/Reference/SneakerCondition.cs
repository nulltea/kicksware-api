using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;
using System.Text.Json.Serialization;
using Newtonsoft.Json.Converters;

namespace Core.Reference
{
	[JsonConverter(typeof(StringEnumConverter))]
	public enum SneakerCondition
	{
		[EnumMember(Value = "New")]
		[Display(Name = "New/Never Worn", ShortName = "New")]
		New,

		[EnumMember(Value = "Gently")]
		[Display(Name = "Gently Used", ShortName = "Gently")]
		GentlyUsed,

		[EnumMember(Value = "Used")]
		[Display(Name = "Used", ShortName = "Used")]
		Used,

		[EnumMember(Value = "Worn")]
		[Display(Name = "Very Worn", ShortName = "Worn")]
		VeryWorn
	}
}