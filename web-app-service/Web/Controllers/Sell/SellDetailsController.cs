using System;
using System.Threading.Tasks;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;
using SmartBreadcrumbs.Attributes;
using Web.Data.Catalog;
using Web.Models;
using Web.Wizards;

namespace Web.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Details", FromAction = "NewProduct", FromController = typeof(SellController))]
		public async Task<ActionResult> SetDetails(string referenceID)
		{
			var service = HttpContext.RequestServices.GetService<ISneakerReferenceService>();
			if (service == null) throw new NotImplementedException($"{nameof(ISneakerReferenceService)} not implemented");

			var reference = await service.FetchUniqueAsync(referenceID);
			if (reference == null) return this.ViewStep(1, new SneakerProductViewModel());

			var sneakerProduct = new SneakerProductViewModel
			{
				ReferenceID = reference.UniqueID,
				ModelSKU = reference.ManufactureSku,
				ModelName = reference.ModelName,
				BrandName = reference.BrandName,
				Description = reference.Description,
				Color = reference.Color,
				Price = reference.Price,
				Size = Catalog.SneakerSizesList[11]
			};

			HeroBreadSubTitle = "Tell us about your sneakers. In details - the more the better";
			return this.ViewStep(1, sneakerProduct);
		}

		[HttpPost]
		public ActionResult Details(SneakerProductViewModel model, bool rollback)
		{
			return rollback ? RedirectToAction("NewProduct") : SetPhotos(model);
		}
	}
}