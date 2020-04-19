﻿using System.IO;
using System.Linq;
using Core.Services;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs;
using SmartBreadcrumbs.Attributes;
using SmartBreadcrumbs.Nodes;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController : Controller
	{
		private readonly ISneakerProductService _service;

		private readonly IWebHostEnvironment _environment;

		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/seller-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Add product listing";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "Search and choose the sneaker model to prefill further details";

		public SellController(ISneakerProductService service, IWebHostEnvironment env) => (_service, _environment) = (service, env);

		[Breadcrumb("Sell", FromAction = "Index", FromController = typeof(HomeController))]
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

		private void AddBreadcrumbNode(string action)
		{
			var baseNode = new MvcBreadcrumbNode("NewProduct", "Sell", "Sell");
			var currentNode = new MvcBreadcrumbNode(action, "Sell", action) {Parent = baseNode};
			ViewData["BreadcrumbNode"] = currentNode;
		}
	}
}