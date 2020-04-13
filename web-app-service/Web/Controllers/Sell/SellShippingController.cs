using System.Linq;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Shipping", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult SetShipping(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "Take care of the shipping. It better be worldwide";
			if (model.ShippingInfo is null || !model.ShippingInfo.Any())
			{
				model.ShippingInfo = Catalog.DefaultShippingInfo;
			}
			AddBreadcrumbNode(nameof(Shipping));
			return this.ViewStep(4, model);
		}

		[HttpPost]
		public ActionResult Shipping(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return SetPayment(model);
			return ShowPreview(model);
		}
	}
}