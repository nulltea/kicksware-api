using System.Linq;
using Core.Entities.References;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;

namespace Web.Controllers
{
	public partial class ShopController
	{
		[HttpGet]
		[Route("shop/brand/{brandID}")]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Brand(string brandID, int page = 1, string sortBy = default)
		{
			var references = InitFilterHandler<SneakerReference>(new {brandID});
			if (!string.IsNullOrEmpty(sortBy)) references.ChooseSortParameter(sortBy);
			references.FetchPage(page);

			var brand = references.FirstOrDefault()?.Brand ?? new SneakerBrand(brandID);

			HeroCoverPath = brand.HeroPath;
			HeroBreadTitle = brand.Name;
			HeroBreadSubTitle = brand.Description;
			HeroLogoPath = brand.LogoPath;

			AddBreadcrumbNode(nameof(Brand), brand.Name);
			return View("References", references);
		}
	}
}