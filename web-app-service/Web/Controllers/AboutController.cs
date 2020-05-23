using System.Collections.Generic;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public class AboutController : Controller
	{
		[Breadcrumb("About", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult About()
		{
			return View(new List<string>
			{
				"/images/heroes/about-hero.jpg",
				"/images/heroes/seller-hero.jpg",
				"/images/heroes/shop-hero.jpg",
				"/images/heroes/seller-hero2.jpg",
				"/images/heroes/contact-hero.jpg",
			});
		}
	}
}