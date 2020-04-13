using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;
using System.Text.Json.Serialization;
using Newtonsoft.Json.Converters;

namespace Core.Reference
{
	[JsonConverter(typeof(StringEnumConverter))]
	public enum SneakerType
	{
		[EnumMember(Value = "Low")]
		[Display(Name = "Low-Top Sneakers", ShortName = "Low-Top")]
		LowTop,

		[EnumMember(Value = "High")]
		[Display(Name = "High-Top Sneakers", ShortName = "High-Top")]
		HighTop,

		[EnumMember(Value = "SlipOns")]
		[Display(Name = "Slip-ons", ShortName = "Slip-ons")]
		SlipOns,

		[EnumMember(Value = "Athletic")]
		[Display(Name = "Athletic Kicks", ShortName = "Athletic")]
		AthleticKicks,

		[EnumMember(Value = "Converse")]
		[Display(Name = "Converse Sneakers", ShortName = "Converse")]
		Converse,

		[EnumMember(Value = "Boots")]
		[Display(Name = "Boots", ShortName = "Boots")]
		Boots,

		[EnumMember(Value = "Formal")]
		[Display(Name = "Formal Shoes", ShortName = "Formal")]
		FormalShoes,

		[EnumMember(Value = "Casual")]
		[Display(Name = "Casual Shoes", ShortName = "Casual")]
		CasualShoes
	}
}