﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace Web.Controllers
{
	public class CheckoutController : Controller
	{
		public IActionResult Checkout()
		{
			return View();
		}
	}
}