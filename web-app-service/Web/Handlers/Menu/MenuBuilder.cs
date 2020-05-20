using System.IO;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;
using Newtonsoft.Json;

namespace Web.Handlers.Menu
{
	public class MenuBuilder<TMenuContent> where TMenuContent : IMenuContent
	{
		private TMenuContent _menuContent;

		public MenuBuilder(string path)
		{
			_menuContent = JsonConvert.DeserializeObject<TMenuContent>(File.ReadAllText(path));
		}

		public IHtmlContent RenderMenu(IHtmlHelper html) => _menuContent.Render(html);
	}
}