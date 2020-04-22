using System.Globalization;

namespace Core.Extension
{
	public static class StringExtension
	{
		public static string ToTitleCase(this string source) => CultureInfo.CurrentCulture.TextInfo.ToTitleCase(source);
	}
}