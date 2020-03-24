using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;

namespace Core.Reference
{
	public enum SearchType
	{
		[EnumMember(Value = "BrandModel")]
		[Display(Name = "Brand & Model", ShortName = "Brand-Model")]
		BrandModel,
		[EnumMember(Value = "ModelId")]
		[Display(Name = "Model Id", ShortName = "ModelId")]
		ModelId,
		[EnumMember(Value = "SKU")]
		[Display(Name = "SKU", ShortName = "SKU")]
		SKU,
	}
}