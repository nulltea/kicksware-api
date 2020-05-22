using System.Linq;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Data.Catalog;
using Web.Models;
using Web.Wizards;

namespace Web.Controllers
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
			return rollback ? SetPayment(model) : ShowPreview(model);
		}
	}
}