using System.Globalization;
using System.Text.RegularExpressions;

namespace Core.Extension
{
	public static class StringExtension
	{
		public static string ToTitleCase(this string source) => CultureInfo.CurrentCulture.TextInfo.ToTitleCase(source);

		public static string ToFormattedID(this string source, string separator = "-") =>
			new Regex("[\\n\\t;,.\\s()\\/]").Replace(source, separator);
	}
}