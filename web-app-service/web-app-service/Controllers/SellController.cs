using Core.Repositories;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;

namespace web_app_service.Controllers
{
	public class SellController : Controller
	{
		private readonly ISneakerProductRepository _repository;

		public SellController(ISneakerProductRepository repository) => _repository = repository;

		// GET: Sell
		public ActionResult AddProduct()
		{
			return View();
		}

		// GET: Sell
		public ActionResult MyProducts()
		{
			return View();
		}

		// POST: Sell/Post
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Post([FromForm]SneakerProductViewModel sneakerProduct)
		{
			var response = _repository.Post(sneakerProduct);

			if (response == null) return Problem();

			return RedirectToAction("Index", "Home");
		}

		// GET: Sell/Edit/5
		public ActionResult Edit(string productId)
		{
			return View();
		}

		// POST: Sell/Edit/5
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Edit(string id, IFormCollection collection)
		{
			try
			{
				// TODO: Add update logic here

				return RedirectToAction(nameof(AddProduct));
			}
			catch
			{
				return View();
			}
		}

		// GET: Sell/Delete/5
		public ActionResult Delete(string productId)
		{
			return View();
		}

		// POST: Sell/Delete/5
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Delete(string productId, IFormCollection collection)
		{
			try
			{
				// TODO: Add delete logic here

				return RedirectToAction(nameof(AddProduct));
			}
			catch
			{
				return View();
			}
		}
	}
}