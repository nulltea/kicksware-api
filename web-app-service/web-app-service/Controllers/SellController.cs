using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace web_app_service.Controllers
{
	public class SellController : Controller
	{
		public static string[] BrandsList =
		{
			"Nike",
			"NikeSB",
			"Air Jordan",
			"Adidas",
			"Reebok",
			"Puma",
			"Asics",
			"Yeezy",
			"Addidas Y-3",
			"Converse",
			"Vans",
			"New Balance",
			"Fila",
			"Bape",
			"Supra",
			"Lacoste",
			"Yeezy",
			"Kappa"
		};

		// GET: Sell
		public ActionResult AddProduct()
		{
			ViewBag.BrandsList = BrandsList;
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
		public ActionResult Edit(int id)
		{
			return View();
		}

		// POST: Sell/Edit/5
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Edit(int id, IFormCollection collection)
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
		public ActionResult Delete(int id)
		{
			return View();
		}

		// POST: Sell/Delete/5
		[HttpPost]
		[ValidateAntiForgeryToken]
		public ActionResult Delete(int id, IFormCollection collection)
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