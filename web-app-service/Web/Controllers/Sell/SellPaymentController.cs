using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;
using Web.Wizards;

namespace Web.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Payment", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult SetPayment(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "So let's talk about the price...";
			model.Price = 500m;
			AddBreadcrumbNode(nameof(Payment));
			return this.ViewStep(3, model);
		}

		[HttpPost]
		public ActionResult Payment(SneakerProductViewModel model, bool rollback)
		{
			return rollback ? SetPhotos(model) : SetShipping(model);
		}
	}
}