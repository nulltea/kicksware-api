using Core.Entities.Products;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;
using Web.Wizards;

namespace Web.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Preview", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult ShowPreview(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "Great job! If you are satisfied with it - send the whole thing to us, and we will publish it to all the sneakerheads around the world";
			AddBreadcrumbNode(nameof(Preview));
			return this.ViewStep(5, model);
		}

		[HttpPost]
		public ActionResult Preview([FromServices] ISneakerProductService service, [FromForm] SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return SetShipping(model);

			var sneakerProduct = model as SneakerProduct;
			var response = service.Store(sneakerProduct);

			if (response == null) return Problem();

			return RedirectToAction("ProductItem", "Shop", new {productId = sneakerProduct.UniqueID});
		}

	}
}