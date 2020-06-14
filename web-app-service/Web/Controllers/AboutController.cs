using System.Collections.Generic;
using Core.Constants;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public class AboutController : Controller
	{
		[Route("about")]
		[Breadcrumb("About", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult About()
		{
			return View(new List<string>
			{
				$"{Constants.FileStoragePath}/heroes/about-hero.jpg",
				$"{Constants.FileStoragePath}/heroes/seller-hero.jpg",
				$"{Constants.FileStoragePath}/heroes/shop-hero.jpg",
				$"{Constants.FileStoragePath}/heroes/seller-hero2.jpg",
				$"{Constants.FileStoragePath}/heroes/contact-hero.jpg",
			});
		}
	}
}