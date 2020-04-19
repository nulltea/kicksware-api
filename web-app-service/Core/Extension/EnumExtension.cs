using System;
using System.Reflection;
using System.Runtime.Serialization;

namespace Core.Extension
{
	public static class EnumExtension
	{
		public static object GetEnumMemberValue(this Enum enumMember)
		{
			var type = enumMember.GetType();
			var field = type.GetField(nameof(enumMember));
			var memberAttr = field.GetCustomAttribute<EnumMemberAttribute>(true);
			return memberAttr.Value;
		}
	}
}