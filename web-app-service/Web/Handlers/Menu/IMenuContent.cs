using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Web.Handlers.Menu
{
	public interface IMenuContent
	{
		string Caption { get; set; }

		string Controller { get; set; }

		string Action { get; set; }

		string RouteValues { get; set; }

		IMenuContent ParentContent { get; set; }

		IHtmlContent Render(IHtmlHelper html);
	}
}