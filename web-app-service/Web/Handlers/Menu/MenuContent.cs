using System.Collections.Generic;
using System.Linq;
using Core.Extension;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Web.Handlers.Menu
{
	public abstract class MenuListContent<TMenuContent> : IMenuContent where TMenuContent: class, IMenuContent
	{
		public string Caption { get; set; }

		public string Controller { get; set; }

		public string Action { get; set; }

		public string RouteValues { get; set; }

		public IMenuContent ParentContent { get; set; }

		public List<TMenuContent> InnerContent { get; set; }

		public abstract IHtmlContent Render(IHtmlHelper html);

		public virtual void FillMissingAttributes()
		{
			if (string.IsNullOrEmpty(RouteValues) && !string.IsNullOrEmpty(Caption)) RouteValues = Caption.ToFormattedID();

			if (InnerContent == null || !InnerContent.Any()) return;

			foreach (var sub in InnerContent)
			{
				if (string.IsNullOrEmpty(sub.Controller)) sub.Controller = Controller;
				if (string.IsNullOrEmpty(sub.Action)) sub.Action = Action;
				sub.ParentContent = this;
			}
		}
	}

	public abstract class MenuContent : IMenuContent
	{
		public string Caption { get; set; }

		public string Controller { get; set; }

		public string Action { get; set; }

		public string RouteValues { get; set; }

		public IMenuContent ParentContent { get; set; }

		public abstract IHtmlContent Render(IHtmlHelper html);

		public virtual void FillMissingAttributes()
		{
			if (string.IsNullOrEmpty(RouteValues) && !string.IsNullOrEmpty(Caption)) RouteValues = Caption.ToFormattedID();
		}
	}
}