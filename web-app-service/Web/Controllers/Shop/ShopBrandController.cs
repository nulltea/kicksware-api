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
		[Route("shop/brand/{brandId}")]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult Brand(string brandId, int page = 1, string sortBy = default)
		{
			var references =
				InitFilterHandler<SneakerReference>(new {brandId}); //TODO custom builder
			if (!string.IsNullOrEmpty(sortBy)) references.ChooseSortParameter(sortBy);
			references.FetchPage(page);
			var brand = references.FirstOrDefault()?.Brand ?? new SneakerBrand(brandId);
			HeroCoverPath = brand.HeroPath;
			HeroBreadTitle = brand.Name;
			HeroBreadSubTitle = brand.Description;
			HeroLogoPath = brand.Logo;

			return View("References", references);
		}
	}
}