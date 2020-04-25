using Core.Entities.Products;
using Core.Entities.References;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public partial class ShopController
	{
		[HttpGet]
		[Route("shop/products")]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Products(string referenceId = default, int page = 1, string sortBy = default)
		{
			var products = InitFilterHandler<SneakerProduct>(new {referenceId});
			if (!string.IsNullOrEmpty(sortBy)) products.ChooseSortParameter(sortBy);
			products.FetchPage(page);
			return View("References", products);
		}
	}
}