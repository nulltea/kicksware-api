﻿using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using web_app_service.Models;

namespace web_app_service.Controllers
{
	public class HomeController : Controller
	{
		public List<HomePageInfo> HomeInfo => new List<HomePageInfo>
		{
			new HomePageInfo
			{
				Title = "Nike ISPA’s Newest Round Of Releases Is Futurist Perfection",
				Image = "home_nike_ispa_envelope.jpeg",
				Description =
					"Nike ISPA continues headstrong into a new era of releases, offering up futuristic design sensibilities in a wrapping of utilitarian purpose. With the past calendar year ripe with innovation though lukewarm in reception, the latest round of their releases is geared to be take things up one more notch with two brand new debuts cut from a technical cloth of their own",
				ButtonCaption = "Shop Now",
				ButtonAction = Url.Action("Products", "Shop")
			},
			new HomePageInfo
			{
				Title = "A Closer Look at adidas Consortium's EVO 4D F&F for Paris Fashion Week",
				Image = "home_addidas_4d.jpg",
				Description =
					"One of the most notable additions to this F&F pair is its blacked-out 4D sole unit — something we’ve only just started to see on other offerings such as the 4D Run 1.0. The insoles have been printed with the exclusive phrase “Consortium EVO 4D Paris Fashion Week 2020,” while the tongue tab sports the familiar Consortium metal eyelet",
				ButtonCaption = "Sell Now",
				ButtonAction = Url.Action("Products", "Shop")
			},
			new HomePageInfo
			{
				Title = "Nike ISPA Armors The Air Max 720 With Rivets From The React Element Soles",
				Image = "home_nike_ispa_720.jpeg",
				Description =
					"With fall and winter’s tempestuous weather right around the corner, it’s important to make sure that your footwear is rugged enough to handle anything Mother Nature might throw at you, so Nike ISPA has prepared a brand-new, highly conceptual Air Max 720 that can get you through even the most arduous conditions. Short for “Improvise, Scavenge, Protect, Adapt,” ISPA...",
				ButtonCaption = "Shop Now",
				ButtonAction = Url.Action("Products", "Shop")
			},
		};

		private readonly ILogger<HomeController> _logger;

		public HomeController(ILogger<HomeController> logger)
		{
			_logger = logger;
		}

		public IActionResult Index()
		{
			ViewBag.FeaturedProducts = ShopController.ProductsList; //TODO
			return View(HomeInfo);
		}

		public IActionResult Privacy()
		{
			return View();
		}

		[ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
		public IActionResult Error()
		{
			return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
		}
	}
}