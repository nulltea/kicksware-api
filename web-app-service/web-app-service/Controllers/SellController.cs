using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using Core.Reference;
using Core.Services;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;

namespace web_app_service.Controllers
{
	public class SellController : Controller
	{
		private readonly ISneakerProductService _service;

		private readonly IWebHostEnvironment _environment;

		public SellController(ISneakerProductService service, IWebHostEnvironment environment) => (_service, _environment) = (service, environment);

		[HttpGet]
		public ActionResult AddProduct()
		{
			var sneakerProduct = new SneakerProductViewModel
			{
				ShippingInfo = Catalog.DefaultShippingInfo
			};
			return View(sneakerProduct);
		}

		[HttpGet]
		public ActionResult MyProducts()
		{
			return View();
		}

		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Store([FromForm]SneakerProductViewModel sneakerProduct)
		{
			foreach (var formFile in sneakerProduct.FormFiles)
			{
				if (formFile.Length <= 0) continue;
				var fileName = Path.ChangeExtension(Path.GetRandomFileName(), Path.GetExtension(formFile.FileName));
				var filePath = Path.Combine(_environment.WebRootPath, "files", fileName);
				
				using var stream = System.IO.File.Create(filePath);
				formFile.CopyTo(stream);
				sneakerProduct.Photos.Add(filePath);
			}
			var response = _service.Store(sneakerProduct);

			if (response == null) return Problem();

			return RedirectToAction("Index", "Home");
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
	}
}