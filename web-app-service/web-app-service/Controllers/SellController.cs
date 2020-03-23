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
	public class SellController : Controller
	{
		private readonly ISneakerProductService _service;

		private readonly IWebHostEnvironment _environment;

		public SellController(ISneakerProductService service,  IWebHostEnvironment env) => (_service, _environment) = (service, env);

		#region New product wizard handle

		public ActionResult NewProduct(SneakerProductViewModel model)
		{
			return this.ViewStep(0, model);
		}

		public ActionResult Search(SneakerProductViewModel model)
		{
			model.Size = Catalog.SneakerSizesList[11];
			return this.ViewStep(1, model);
		}

		public ActionResult Details(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(0, model);

			return this.ViewStep(2, model);
		}

		public ActionResult Photos(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(1, model);
			if (model.FormFiles != null && model.FormFiles.Any())
			{
				foreach (var formFile in model.FormFiles)
				{
					if (formFile.Length <= 0) continue;
					var fileName = Path.ChangeExtension(Path.GetRandomFileName(), Path.GetExtension(formFile.FileName));
					var filePath = Path.Combine(_environment.WebRootPath, "files", fileName);

					using var stream = System.IO.File.Create(filePath);
					formFile.CopyTo(stream);
					model.Photos.Add(filePath);
				}
			}

			model.Price = 500m;

			return this.ViewStep(3, model);
		}

		public ActionResult Payment(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(2, model);
			if (model.ShippingInfo is null || !model.ShippingInfo.Any())
			{
				model.ShippingInfo = Catalog.DefaultShippingInfo;
			}
			return this.ViewStep(4, model);
		}

		public ActionResult Shipping(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(3, model);
			return this.ViewStep(5, model);
		}

		public ActionResult Preview(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(4, model);
			return RedirectToAction("Index", "Home");
		}

		#endregion
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