using Core.Entities.References;
using Core.Services;
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
		[Breadcrumb("Sell", FromAction = "Index", FromController = typeof(HomeController))]
		public JsonResult SearchAuto([FromServices] IReferenceSearchService service, [FromQuery] string prefix)
		{
			var references = service.Search(prefix);
			return Json(references);
		}

		[HttpPost]
		public ActionResult Search(SneakerReference reference)
		{
			if (reference is null || string.IsNullOrWhiteSpace(reference.UniqueID)) return this.ViewStep(1, new SneakerProductViewModel());

;			var sneakerProduct = new SneakerProductViewModel
			{
				ReferenceID = reference.UniqueID,
				ModelSKU = reference.ManufactureSku,
				ModelName = reference.ModelName,
				BrandName = reference.BrandName,
				Description = reference.Description,
				Color = reference.Color,
				Price = reference.Price,
				Size = Catalog.SneakerSizesList[11]
			};

			return new JsonResult(new {redirectUrl = Url.Action("SetDetails", sneakerProduct)});
		}
	}
}