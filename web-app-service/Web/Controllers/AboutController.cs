using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace web_app_service.Controllers
{
	public class AboutController : Controller
	{
		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/about-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "About us";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "Get to know us better";

		[Breadcrumb("About", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult About()
		{
			return View();
		}
	}
}