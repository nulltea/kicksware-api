using System;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;
using Web.Models;
using Web.Wizards;

namespace Web.Controllers
{
	public partial class SellController
	{
		[HttpGet]
		[Breadcrumb("Photos", FromAction = "NewProduct", FromController = typeof(SellController))]
		public ActionResult SetPhotos(SneakerProductViewModel model)
		{
			HeroBreadSubTitle = "Show buyers the best in your product";
			AddBreadcrumbNode(nameof(Photos));
			return this.ViewStep(2, model);
		}

		[HttpPost]
		public async Task<ActionResult> Photos(SneakerProductViewModel model, bool rollback)
		{
			if (model == null) throw new ArgumentNullException(nameof(model));

			if (rollback) return await SetDetails(model.ReferenceID);
			if (model.FormFiles != null && model.FormFiles.Any())
			{
				foreach (var formFile in model.FormFiles)
				{
					if (formFile.Length <= 0) continue;
					var fileName = Path.ChangeExtension(Path.GetRandomFileName(), Path.GetExtension(formFile.FileName));
					var filePath = Path.Combine(_environment.WebRootPath, "files/photos/products", fileName);

					using var stream = System.IO.File.Create(filePath);
					formFile.CopyTo(stream);
					model.Photos.Add(filePath);
				}
			}

			return SetPayment(model);
		}
	}
}