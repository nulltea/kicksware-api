using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace web_app_service.Controllers
{
	public class ContactController : Controller
	{
		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/contact-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Get in touch with us";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "We are thrilled to meet with you";

		[Breadcrumb("Contact", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Contact()
		{
			return View();
		}
	}
}