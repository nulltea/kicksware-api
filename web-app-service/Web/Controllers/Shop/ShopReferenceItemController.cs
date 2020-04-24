using System.Collections.Generic;
using Core.Entities.References;
using Core.Gateway;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public partial class ShopController
	{
		[HttpGet]
		[Route("shop/references/{referenceId}")]
		[Breadcrumb("Product item", FromAction = "References", FromController = typeof(ShopController))]
		public IActionResult ReferenceItem(string referenceId, [FromServices] ISneakerReferenceService service)
		{
			var reference = service.FetchUnique(referenceId);

			if (reference is null) return NotFound();

			ViewBag.RelatedReferences = service.GetRelated(reference, new RequestParams {TakeCount = 12});

			return View(reference);
		}
	}
}