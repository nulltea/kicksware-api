using System.ComponentModel.DataAnnotations;

namespace Core.Reference
{
	public enum SneakerCondition
	{
		[Display(Name = "New/Never Worn", ShortName = "New")]
		New,

		[Display(Name = "Gently Used", ShortName = "Gently")]
		GentlyUsed,

		[Display(Name = "Used", ShortName = "Used")]
		Used,

		[Display(Name = "Very Worn", ShortName = "Worn")]
		VeryWorn
	}
}