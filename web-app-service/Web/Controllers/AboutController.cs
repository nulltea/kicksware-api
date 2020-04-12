using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace web_app_service.Controllers
{
    public class AboutController : Controller
    {
        [Breadcrumb("About", FromAction = "Index", FromController = typeof(HomeController))]
        public IActionResult About()
        {
            return View();
        }
    }
}