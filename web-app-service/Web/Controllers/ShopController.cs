using System.Collections.Generic;
using System.Linq;
using Core.Entities.Reference;
using Core.Model.Parameters;
using Core.Reference;
using Core.Services;
using Infrastructure.Usecase.Models;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using web_app_service.Data.Reference_Data;

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
			var references = new FilteredModel<SneakerReference>(_service, page);
			references.AddGroup("Brand", "brandname")
				.AssignParameters(Catalog.SneakerBrandsList);
			references.AddGroup("Size", "size").AssignParameters(Catalog.SneakerSizesList,
				size => new FilterParameter(size.Europe.ToString("#.#"), size)); //TODO foreign constraint filter
			references.AddGroup("Color", "color", ExpressionType.Or).AssignParameters(Catalog.FilterColors,
				color => new FilterParameter(color.Name, color.Name, ExpressionType.Like) {SourceValue = color});
			references.AddGroup("Price", "price", ExpressionType.Between).AssignParameters(
				new FilterParameter("Price (min)", 100) {Checked = true},
				new FilterParameter("Price (max)", 1000) {Checked = true});
			references.AddGroup("Condition", "condition").AssignParameters(typeof(SneakerCondition));

			references.FetchPage(page);
			return View(references);
		}

		[HttpPost]
		public IActionResult RequestFilter()
		{
			return View("Products");
		}

		[HttpGet]
		[Breadcrumb("Product item", FromAction = "Products", FromController = typeof(ShopController))]
		public IActionResult ProductItem(string productId)
		{
			var product = _service.FetchUnique(productId);

			if (product == null) return NotFound();
			//ViewBag.RelatedProducts = ProductsList; //TODO search related
			return View(product);
		}
	}
}