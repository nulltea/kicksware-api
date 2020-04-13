using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;

namespace Core.Reference
{
	public enum SortType
	{
		[EnumMember(Value = "Popular")]
		[Display(Name = "Popular", ShortName = "Popular")]
		Popular,

		[EnumMember(Value = "Newest")]
		[Display(Name = "Newest", ShortName = "New")]
		Newest,

		[EnumMember(Value = "Featured")]
		[Display(Name = "Featured", ShortName = "Featured")]
		Featured,

		[EnumMember(Value = "PriceFromLow")]
		[Display(Name = "Price (Low-High)", ShortName = "Price Low")]
		PriceFromLow,

		[EnumMember(Value = "PriceFromHigh")]
		[Display(Name = "Price (High-Low)", ShortName = "Price High")]
		PriceFromHigh
	}
}