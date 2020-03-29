using System.Linq;
using Core.Entities.Reference;
using Core.Reference;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		public JsonResult SearchAuto([FromServices] IReferenceSearchService service, [FromQuery]string prefix)
		{
			var references = service.Search(prefix);
			return Json(references.ToArray());
		}

		public ActionResult Search(SneakerReference reference)
		{
			var sneakerProduct = new SneakerProductViewModel
			{
				ModelRefId = reference.UniqueId,
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