using System.Runtime.Serialization;

namespace Core.Reference
{
	public enum ExpressionType
	{
		[EnumMember(Value = "$and")]
		And,

		[EnumMember(Value = "$or")]
		Or,

		[EnumMember(Value = "$in")]
		In,

		[EnumMember(Value = "$nin")]
		NotIn,

		[EnumMember(Value = "$bt")]
		Between,

		[EnumMember(Value = "$eq")]
		Equal,

		[EnumMember(Value = "$ne")]
		NotEqual,

		[EnumMember(Value = "$text")]
		Like,

		[EnumMember(Value = "$regex")]
		LikeRegex,

		[EnumMember(Value = "$ge")]
		GreaterThen,

		[EnumMember(Value = "$gte")]
		GreaterThanOrEqual,

		[EnumMember(Value = "$le")]
		LessThan,

		[EnumMember(Value = "$lte")]
		LessThanOrEqual
	}
}