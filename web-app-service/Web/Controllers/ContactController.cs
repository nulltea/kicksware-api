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
		[Breadcrumb("Contact", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Contact()
		{
			return View();
		}
	}
}