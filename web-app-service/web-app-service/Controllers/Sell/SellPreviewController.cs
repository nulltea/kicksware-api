using Core.Entities.Products;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		public ActionResult Preview([FromServices] ISneakerProductService service, [FromForm] SneakerProductViewModel model)
		{
			var sneakerProduct = model as SneakerProduct;
			var response = service.Store(sneakerProduct);

			if (response == null) return Problem();

			return RedirectToAction("Index", "Home");
		}

	}
}