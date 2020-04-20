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
			var field = type.GetField(enumMember.ToString());
			var memberAttr = field.GetCustomAttribute<EnumMemberAttribute>(true);
			return memberAttr.Value;
		}

		public static T GetEnumAttribute<T>(this Enum enumMember) where T : Attribute
		{
			var type = enumMember.GetType();
			var field = type.GetField(enumMember.ToString());
			var attr = field.GetCustomAttribute<T>(true);
			return attr;
		}
	}
}