using System.Reflection;
using System.Runtime.Serialization;

namespace Core.Reference
{
	public enum Gender
	{
		[EnumMember(Value = "")]
		None,

		[EnumMember(Value = "Mens")]
		Mens,

		[EnumMember(Value = "Womens")]
		Womens,

		[EnumMember(Value = "Unisex")]
		Unisex,

		[EnumMember(Value = "Youth")]
		Youth,

		[EnumMember(Value = "Preschool")]
		Preschool,

		[EnumMember(Value = "Toddler")]
		Toddler
	}
}