using System.Collections.Generic;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Web.Handlers.Menu
{
	public abstract class MenuListContent<TMenuContent> : IMenuContent where TMenuContent: IMenuContent
	{
		public string Caption { get; set; }

		public string Controller { get; set; }

		public string Action { get; set; }

		public object RouteValues { get; set; }

		public List<TMenuContent> InnerContent { get; set; }

		public abstract IHtmlContent Render(IHtmlHelper html);
	}

	public abstract class MenuContent : IMenuContent
	{
		public string Caption { get; set; }

		public string Controller { get; set; }

		public string Action { get; set; }

		public object RouteValues { get; set; }

		public abstract IHtmlContent Render(IHtmlHelper html);

	}
}