using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Sell", FromAction = "Index", FromController = typeof(HomeController))]
		public ActionResult SetDetails(SneakerProductViewModel model)
		{
			return this.ViewStep(1, model);
		}

		[Breadcrumb("Sell", FromAction = "Index", FromController = typeof(HomeController))]
		public ActionResult Details(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(0, model);

			return this.ViewStep(2, model);
		}
	}
}