using System.Collections.Generic;
using System.Linq;
using Core.Extension;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Web.Handlers.Menu
{
	/// <li class="expandable" id="shop-menu">
	/// 	<a asp-controller="Shop" asp-action="References">SHOP</a>
	/// 	<div class="sub-panel">
	/// 		<ul>
	/// 			...
	/// 		</ul>
	/// 	</div>
	/// </li>
	public class ShopMenuContent : MenuListContent<BrandSubmenuContent>
	{
		public override IHtmlContent Render(IHtmlHelper html)
		{
			FillMissingAttributes();

			var subPanel = new TagBuilder("div");
			subPanel.AddCssClass("sub-panel");

			var brandsMenu = new TagBuilder("ul");

			foreach (var brand in InnerContent)
			{
				brandsMenu.InnerHtml.AppendHtml(brand.Render(html));
			}
			subPanel.InnerHtml.AppendHtml(brandsMenu);

			return subPanel;
		}
	}

	/// <li class="expandable" id="brand-models">
	/// 	<a asp-controller="Action" asp-action="Contoller">Caption</a>
	/// 	<div class="sub-panel">
	/// 		...
	/// 	</div>
	/// </li>
	public class BrandSubmenuContent : MenuListContent<BrandSubgroupContent>
	{
		public override IHtmlContent Render(IHtmlHelper html)
		{
			FillMissingAttributes();

			var expandableItem = new TagBuilder("li");
			expandableItem.AddCssClass("expandable");
			expandableItem.Attributes["id"] = $"{Caption.ToLower()}-models";

			var link = html.ActionLink(Caption.ToUpper(), Action, Controller, new {brandID = RouteValues});
			expandableItem.InnerHtml.AppendHtml(link);

			if (InnerContent is null) return expandableItem;

			var subPanel = new TagBuilder("div");
			subPanel.AddCssClass("sub-panel");

			if (InnerContent.Count == 1)
			{
				InnerContent.First().FillMissingAttributes();
				foreach (var model in InnerContent.First().InnerContent)
				{
					subPanel.InnerHtml.AppendHtml(model.Render(html));
				}
			}
			else
			{
				foreach (var group in InnerContent)
				{
					subPanel.InnerHtml.AppendHtml(group.Render(html));
				}
			}

			expandableItem.InnerHtml.AppendHtml(subPanel);

			return expandableItem;
		}

		public override void FillMissingAttributes()
		{
			if (string.IsNullOrEmpty(RouteValues) && !string.IsNullOrEmpty(Caption)) RouteValues = Caption.ToFormattedID(" "); // TODO
			base.FillMissingAttributes();
		}
	}

	/// <div class="sub-group">
	/// 	<h3>Caption</h3>
	/// 	...
	/// </div>
	public class BrandSubgroupContent : MenuListContent<ModelSubmenuContent>
	{
		public override IHtmlContent Render(IHtmlHelper html)
		{
			FillMissingAttributes();

			var subGroup = new TagBuilder("div");
			subGroup.AddCssClass("sub-group");

			if (!string.IsNullOrEmpty(Caption))
			{
				var caption = new TagBuilder("h3");
				caption.InnerHtml.Append(Caption.ToUpper());
				subGroup.InnerHtml.AppendHtml(caption);
			}


			foreach (var model in InnerContent)
			{
				subGroup.InnerHtml.AppendHtml(model.Render(html));
			}

			return subGroup;
		}

		public override void FillMissingAttributes()
		{
			foreach (var sub in InnerContent)
			{
				if (string.IsNullOrEmpty(sub.Controller)) sub.Controller = Controller;
				if (string.IsNullOrEmpty(sub.Action)) sub.Action = Action;
				sub.ParentContent = ParentContent ?? this;
			}
		}
	}

	/// <div class="sub-item">
	/// 	<a asp-controller="controller" asp-action="action">Caption</a>
	/// </div>
	public class ModelSubmenuContent : MenuContent
	{
		public override IHtmlContent Render(IHtmlHelper html)
		{
			FillMissingAttributes();

			var subItem = new TagBuilder("div");
			subItem.AddCssClass("sub-item");

			object routeValues = new {modelID = RouteValues};

			if (Action == "Brand")
			{
				routeValues = new {brandID = RouteValues};
			}

			var link = html.ActionLink(Caption, Action, Controller, routeValues);
			subItem.InnerHtml.AppendHtml(link);

			return subItem;
		}

		public override void FillMissingAttributes()
		{
			if (string.IsNullOrEmpty(RouteValues) && !string.IsNullOrEmpty(Caption))
			{
				if (!string.IsNullOrEmpty(ParentContent?.RouteValues) && !new[]{"More"}.Contains(ParentContent?.RouteValues))
				{
					RouteValues = string.Join("_", ParentContent?.RouteValues, Caption.ToFormattedID());
					return;
				}

				RouteValues = Caption.ToFormattedID(Action == "Brand" ? " " : "-");
			}
		}
	}


}