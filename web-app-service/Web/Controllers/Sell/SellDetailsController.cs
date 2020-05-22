using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;
using Web.Wizards;

namespace Web.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Details", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult SetDetails(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "Tell us about your sneakers. In details - the more the better";

			return this.ViewStep(1, model ?? new SneakerProductViewModel());
		}

		[HttpPost]
		public ActionResult Details(SneakerProductViewModel model, bool rollback)
		{
			return rollback ? RedirectToAction("NewProduct") : SetPhotos(model);
		}
	}
}