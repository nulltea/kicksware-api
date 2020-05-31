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

		public static T GetEnumByMemberValue<T>(this string memberValue) where T : Enum
		{
			var values = (T[]) Enum.GetValues(typeof(T));
			foreach (var enumValue in values)
			{
				if (enumValue.ToString().ToUpper().Equals(memberValue.ToUpper())) return enumValue;
			}

			return values[0];
		}
	}
}