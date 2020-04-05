using System.IO;
using Core.Constants;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace web_app_service.Utils.Helpers
{
	public static partial class CustomHelpers
	{
		public static IHtmlContent VectorIconRender(this IHtmlHelper helper, string icon, params object[] attr)
		{
			return new HtmlFormattableString(File.ReadAllText(Path.Combine(Constants.ImagesPath, icon)), attr);
		}
	}
}