using Core.Entities.References;
using Core.Gateway;
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
			var references = service.Search(prefix, new RequestParams{Limit = 12});
			return Json(references);
		}

		[HttpPost]
		public ActionResult Search([FromBody] string referenceID)
		{
			return new JsonResult(new {redirectUrl = Url.Action("SetDetails"), }); //TODO SPA
		}
	}
}