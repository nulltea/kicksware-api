using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;

namespace Core.Extension
{
	public static class TypeExtension
	{
		public static Dictionary<string, object> ToMap(this object source)
		{
			return source.GetType().GetProperties(BindingFlags.Instance | BindingFlags.Public)
				.ToDictionary(prop => prop.Name, prop => prop.GetValue(source, null));
		}
	}
}