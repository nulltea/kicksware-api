using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Core.Constants;
using Core.Entities.References;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Utils;

namespace Web.Controllers
{
	public partial class ShopController
	{
		[HttpGet]
		[Route("shop/products/{productId}")]
		[Breadcrumb("Product item", FromAction = "References", FromController = typeof(ShopController))]
		public async Task<IActionResult> ProductItem(string productId, [FromServices]ISneakerProductService service, [FromServices] ISneakerReferenceService references)
		{
			var product = await service.FetchUniqueAsync(productId);

			if (product is null) return NotFound();

			//ViewBag.RelatedProducts = ProductsList; //TODO search related

			var reference = await references.FetchUniqueAsync(product.ReferenceID);
			var brand = reference?.Brand ?? product.Brand;
			var baseModel = reference?.BaseModel ?? product.Model;

			HeroCoverPath = baseModel.HeroPath ?? brand.HeroPath;
			HeroBreadTitle = product.ModelName;
			HeroBreadSubTitle = brand.Name;
			HeroLogoPath = brand.LogoPath;

			AddBreadcrumbNode(nameof(ProductItem), product.ModelName);

			return View(product.ToViewModel());
		}
	}
}