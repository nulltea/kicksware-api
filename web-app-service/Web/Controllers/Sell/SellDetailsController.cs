using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Details", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult SetDetails(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "Tell us about your sneakers. In details - the more the better";

			return this.ViewStep(1, model);
		}

		[HttpPost]
		public ActionResult Details(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return RedirectToAction("NewProduct");

			return SetPhotos(model);
		}
	}
}