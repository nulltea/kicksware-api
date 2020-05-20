using System.Collections.Generic;
using System.Linq;
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
			//var expandableItem = new TagBuilder("li");
			//expandableItem.AddCssClass("expandable");
			//expandableItem.Attributes["id"] = $"{Caption.ToLower()}-models";

			//var link = html.ActionLink(Caption, Action, Controller, RouteValues);
			//expandableItem.InnerHtml.AppendHtml(link);

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
			var expandableItem = new TagBuilder("li");
			expandableItem.AddCssClass("expandable");
			expandableItem.Attributes["id"] = $"{Caption.ToLower()}-models";

			var link = html.ActionLink(Caption.ToUpper(), Action, Controller, RouteValues);
			expandableItem.InnerHtml.AppendHtml(link);

			if (InnerContent is null) return expandableItem;

			var subPanel = new TagBuilder("div");
			subPanel.AddCssClass("sub-panel");

			if (InnerContent.Count == 1)
			{
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
	}

	/// <div class="sub-group">
	/// 	<h3>Caption</h3>
	/// 	...
	/// </div>
	public class BrandSubgroupContent : MenuListContent<ModelSubmenuContent>
	{
		public override IHtmlContent Render(IHtmlHelper html)
		{
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
	}

	/// <div class="sub-item">
	/// 	<a asp-controller="controller" asp-action="action">Caption</a>
	/// </div>
	public class ModelSubmenuContent : MenuContent
	{
		public override IHtmlContent Render(IHtmlHelper html)
		{
			var subItem = new TagBuilder("div");
			subItem.AddCssClass("sub-item");

			var link = html.ActionLink(Caption, Action, Controller, RouteValues);
			subItem.InnerHtml.AppendHtml(link);

			return subItem;
		}
	}
}