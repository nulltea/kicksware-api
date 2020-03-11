using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Threading.Tasks;
using Core.Repositories;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Data.Reference_Data;

namespace web_app_service.Controllers
{
	public class SellController : Controller
	{
		public SellController(ISneakerProductRepository repository) {}

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
		public ActionResult Post(IFormCollection collection)
		{
			try
			{
				// TODO: Add insert logic here

				return RedirectToAction(nameof(AddProduct));
			}
			catch
			{
				return View();
			}
		}

		// GET: Sell/Edit/5
		public ActionResult Edit(string id)
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
		public ActionResult Delete(string id)
		{
			return View();
		}

		// POST: Sell/Delete/5
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Delete(string id, IFormCollection collection)
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