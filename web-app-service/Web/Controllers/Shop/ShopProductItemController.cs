using Core.Services;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public partial class ShopController
	{
		[HttpGet]
		[Route("shop/products/{productId}")]
		[Breadcrumb("Product item", FromAction = "References", FromController = typeof(ShopController))]
		public IActionResult ProductItem(string productId, [FromServices]ISneakerProductService service)
		{
			var product = service.FetchUnique(productId);

			if (product == null) return NotFound();
			//ViewBag.RelatedProducts = ProductsList; //TODO search related
			return View();
		}
	}
}