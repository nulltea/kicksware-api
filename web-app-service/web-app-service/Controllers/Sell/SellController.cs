using System.IO;
using System.Linq;
using Core.Services;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController : Controller
	{
		private readonly ISneakerProductService _service;

		private readonly IWebHostEnvironment _environment;

		public SellController(ISneakerProductService service, IWebHostEnvironment env) => (_service, _environment) = (service, env);

		public ActionResult NewProduct([FromServices] IReferenceSearchService service)
		{
			//todo set most relevant references
			return this.ViewStep(0, null);
		}

		[HttpGet]
		public ActionResult MyProducts()
		{
			return View();
		}

		[HttpGet]
		public ActionResult ModifyProduct(string productId)
		{
			return View();
		}

		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Modify([FromForm]SneakerProductViewModel sneakerProduct)
		{
			if (!_service.Modify(sneakerProduct)) return Problem();


			return RedirectToAction("Index", "Home");
		}

		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Delete(string productId)
		{
			if (!_service.Remove(productId)) return Problem();

			return RedirectToAction("Index", "Home");
		}

		//foreach (var formFile in sneakerProduct.FormFiles)
		//{
		//	if (formFile.Length <= 0) continue;
		//	var fileName = Path.ChangeExtension(Path.GetRandomFileName(), Path.GetExtension(formFile.FileName));
		//	var filePath = Path.Combine(_environment.WebRootPath, "files", fileName);

		//	using var stream = System.IO.File.Create(filePath);
		//	formFile.CopyTo(stream);
		//	sneakerProduct.Photos.Add(filePath);
		//}
		//var response = _service.Store(sneakerProduct);

		//if (response == null) return Problem();

		//return RedirectToAction("Index", "Home");
	}
}