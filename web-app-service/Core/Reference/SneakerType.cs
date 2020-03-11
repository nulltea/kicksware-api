using System.ComponentModel.DataAnnotations;

namespace Core.Reference
{
	public enum SneakerType
	{
		[Display(Name = "Low-Top Sneakers", ShortName = "Low-Top")]
		LowTop,

		[Display(Name = "High-Top Sneakers", ShortName = "High-Top")]
		HighTop,

		[Display(Name = "Slip-ons", ShortName = "Slip-ons")]
		SlipOns,

		[Display(Name = "Athletic Kicks", ShortName = "Athletic")]
		AthleticKicks,

		[Display(Name = "Converse Sneakers", ShortName = "Converse")]
		Converse,

		[Display(Name = "Boots", ShortName = "Boots")]
		Boots,

		[Display(Name = "Formal Shoes", ShortName = "Formal")]
		FormalShoes,

		[Display(Name = "Casual Shoes", ShortName = "Casual")]
		CasualShoes
	}
}