using System;
using System.Reflection;
using System.Runtime.Serialization;

namespace Core.Extension
{
	public static class EnumExtension
	{
		public static string GetEnumMemberValue(this Enum source)
		{
			var type = source.GetType();
			var field = type.GetField(source.ToString());
			var memberAttr = field?.GetCustomAttribute<EnumMemberAttribute>(true);
			return memberAttr?.Value;
		}

		public static T GetEnumAttribute<T>(this Enum source) where T : Attribute
		{
			var type = source.GetType();
			var field = type.GetField(source.ToString());
			var attr = field.GetCustomAttribute<T>(true);
			return attr;
		}
	}
}