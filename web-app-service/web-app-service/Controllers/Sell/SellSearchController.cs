using System.Linq;
using Core.Entities.Reference;
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

		public ActionResult Search(SneakerReference model)
		{
			var sneakerProduct = new SneakerProductViewModel {Size = Catalog.SneakerSizesList[11]};
			return this.ViewStep(1, sneakerProduct);
		}
	}
}