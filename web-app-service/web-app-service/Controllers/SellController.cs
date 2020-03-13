using Core.Services;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;

namespace web_app_service.Controllers
{
	public class SellController : Controller
	{
		private readonly ISneakerProductService _service;

		public SellController(ISneakerProductService service) => _service = service;

		[HttpGet]
		public ActionResult AddProduct()
		{
			return View();
		}

		[HttpGet]
		public ActionResult MyProducts()
		{
			return View();
		}

		// POST: Sell/Post
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Post([FromForm]SneakerProductViewModel sneakerProduct)
		{
			var response = _service.Store(sneakerProduct);

			if (response == null) return Problem();

			return RedirectToAction("Index", "Home");
		}

		[HttpGet]
		public ActionResult Modify(string productId)
		{
			return View();
		}

		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Modify([FromForm]SneakerProductViewModel sneakerProduct)
		{
			var response = _service.Store(sneakerProduct);

			if (response == null) return Problem();

			return RedirectToAction("Index", "Home");
		}


		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Delete(string productId)
		{
			if (!_service.Remove(productId)) return Problem();

			return RedirectToAction("Index", "Home");
		}
	}
}