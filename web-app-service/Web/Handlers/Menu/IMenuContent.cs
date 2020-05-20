using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Web.Handlers.Menu
{
	public interface IMenuContent
	{
		string Caption { get; set; }

		string Controller { get; set; }

		string Action { get; set; }

		object RouteValues { get; set; }

		IHtmlContent Render(IHtmlHelper html);
	}
}