using System.Runtime.Serialization;

namespace Core.Entities.Reference
{
	public enum Gender
	{
		[EnumMember(Value = "Mens")]
		Mens,

		[EnumMember(Value = "Womens")]
		Womens,

		[EnumMember(Value = "Unisex")]
		Unisex
	}
}