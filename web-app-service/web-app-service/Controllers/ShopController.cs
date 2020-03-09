using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;

namespace web_app_service.Controllers
{
	public class ShopController : Controller
	{
		public static List<SneakerProduct> ProductsList => new List<SneakerProduct>
		{
			new SneakerProduct
			{
				Id = "1",
				BrandName = "Nike",
				ModelName = "Air Max 720 Ispa",
				Description = "Crazy ones love them AF. My gf gave me these, I ll never sell them!",
				Price = 999m,
				ConditionIndex = 99.9m,
				Images = {"Nike-ISPA-Air-Max-720.jpg"},
				AddedAt = new DateTime(2020, 2, 6)
			},
			new SneakerProduct
			{
				Id = "2",
				BrandName = "Nike",
				ModelName = "Jordan OG",
				Description = "Nike Jordan OG",
				Price = 200m,
				ConditionIndex = 89m,
				Images = {"Nike_Jordan_OG.jpg"},
				AddedAt = new DateTime(2019, 12, 30)
			},
			new SneakerProduct
			{
				Id = "3",
				BrandName = "Nike",
				ModelName = "Air Force 1",
				Description = "Nike Air Force 1 Blue Wight Green",
				Price = 150m,
				ConditionIndex = 77m,
				Images = {"Nike-Af1-Blue-Wight-Green.jpg"},
				AddedAt = new DateTime(2020, 2, 1)
			},
			new SneakerProduct
			{
				Id = "4",
				BrandName = "Nike",
				ModelName = "Jordan Proto Max",
				Description = "Nike Jordan Proto Max 720",
				Price = 240m,
				ConditionIndex = 100m,
				Images = {"Nike-Jordan-Proto-Max-720.jpg"},
				AddedAt = new DateTime(2020, 1, 7)
			},
		};

		public IActionResult Products()
		{
			return View(ProductsList);
		}

		public IActionResult ProductItem(string productId)
		{
			var product = ProductsList.FirstOrDefault(p => p.Id == productId);
			if (product == null) return null;
			ViewBag.RelatedProducts = ProductsList;
			return View(product);
		}
	}
}