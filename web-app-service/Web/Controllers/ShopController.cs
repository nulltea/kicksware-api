using System.Linq;
using Core.Entities.Reference;
using Core.Services;
using Infrastructure.Usecase.Models;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace web_app_service.Controllers
{
	public class ShopController : Controller
	{
		private readonly ISneakerReferenceService _service;

		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/shop-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Buy sneakers";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "Select and buy whatever kicks you like";

		public ShopController(ISneakerReferenceService service) => _service = service;

		[HttpGet]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Products(int page = 1)
		{
			var products = new PagedModelList<SneakerReference>(_service, page);
			return View(products.FetchPage(page));
		}

		[HttpGet]
		[Breadcrumb("Product item", FromAction = "Products", FromController = typeof(ShopController))]
		public IActionResult ProductItem(string productId)
		{
			var product = _service.FetchOne(productId);

			if (product == null) return NotFound();
			//ViewBag.RelatedProducts = ProductsList; //TODO search related
			return View(product);
		}
	}
}