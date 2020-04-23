using System.Collections.Generic;
using System.IO;
using System.Linq;
using Core.Constants;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;

namespace Web.Utils.Helpers
{
	public static partial class CustomHelpers
	{
		public static IHtmlContent VectorIconRender(this IHtmlHelper helper, string icon)
		{
			var iconNode = HtmlAgilityPack.HtmlNode.CreateNode(File.ReadAllText(Path.Combine(Constants.ImagesPath, icon)));
			return new HtmlString(iconNode.OuterHtml);
		}

		public static IHtmlContent VectorIconRender(this IHtmlHelper helper, string icon, params string[] classes)
		{
			var iconNode = HtmlAgilityPack.HtmlNode.CreateNode(File.ReadAllText(Path.Combine(Constants.ImagesPath, icon)));
			classes.ToList().ForEach(iconNode.AddClass);
			return new HtmlString(iconNode.OuterHtml);
		}
	}
}