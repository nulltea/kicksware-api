using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.Reference;
using Core.Model;
using Core.Model.Parameters;
using Core.Reference;
using Core.Services;
using Infrastructure.Usecase.Models;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Abstractions;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;
using Microsoft.AspNetCore.Routing;
using Newtonsoft.Json;
using SmartBreadcrumbs.Attributes;
using web_app_service.Data.Reference_Data;
using web_app_service.Utils.Extensions;

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
		public IActionResult Products(int page = 1, string sortBy = default)
		{
			var references = InitFilterHandler(_service);
			if (!string.IsNullOrEmpty(sortBy)) references.ChooseSortParameter(sortBy);
			references.FetchPage(page);
			return View(references);
		}

		public async Task<IActionResult> RequestFilter(Dictionary<string, Dictionary<bool, string>> filterInputs, string sortBy = default)
		{
			var references = InitFilterHandler(_service);
			if (filterInputs != null && filterInputs.Any())
			{
				references.ApplyUserInputs(filterInputs.ToDictionary(
					input => input.Key,
					input => (input.Value.First().Key, JsonConvert.DeserializeObject(input.Value.First().Value))
				));
			}
			if (!string.IsNullOrEmpty(sortBy)) references.ChooseSortParameter(sortBy);

			references.FetchFiltered();

			return Json(new
			{
				content = await this.RenderViewAsync("ReferencesPartial", references, true),
				length = references.CountTotal,
				pageSize = references.PageSize
			});
		}

		private IFilteredModel<SneakerReference> InitFilterHandler(ICommonService<SneakerReference> service)
		{
			var referenceHandler = new FilteredModel<SneakerReference>(service);

			// Setup filter groups & options
			referenceHandler.AddFilterGroup("Brand", "brandname")
				.AssignParameters(Catalog.SneakerBrandsList);
			referenceHandler.AddForeignFilterGroup<SneakerProduct>("Size", "size")
				.AssignParameters(Catalog.SneakerSizesList, size =>
					new FilterParameter(size.Europe.ToString("#.#"), size));
			referenceHandler.AddFilterGroup("Color", "color", ExpressionType.Or)
				.AssignParameters(Catalog.FilterColors, color =>
					new FilterParameter(color.Name, color.Name.ToUpper(), ExpressionType.Regex) {SourceValue = color});
			referenceHandler.AddFilterGroup("Price", "price", ExpressionType.And)
				.AssignParameters(
					new FilterParameter("Price (min)", 0, ExpressionType.GreaterThanOrEqual)
						{Checked = true, ImmutableValue = false},
					new FilterParameter("Price (max)", 1000, ExpressionType.LessThanOrEqual)
						{Checked = true, ImmutableValue = false}
			);
			referenceHandler.AddForeignFilterGroup<SneakerProduct>("Condition", "condition")
				.AssignParameters(typeof(SneakerCondition));

			// Setup sort criteria
			referenceHandler.AddSortParameters(criterion => criterion switch
			{
				SortCriteria.Popular => new SortParameter(criterion, "likes"),
				SortCriteria.Newest => new SortParameter(criterion, "released"),
				SortCriteria.Featured => new SortParameter(criterion, "likes"),
				SortCriteria.PriceFromLow => new SortParameter(criterion, "price", SortDirection.Ascending),
				SortCriteria.PriceFromHigh => new SortParameter(criterion, "price"),
				_ => throw new ArgumentException("No such sort criteria")
			});

			return referenceHandler;
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

	public class FilterInput
	{
		public bool Checked { get; set; }

		[JsonProperty(PropertyName="value")]
		public object Value { get; set; }
	}
}