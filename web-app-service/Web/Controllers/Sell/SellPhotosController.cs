using System.IO;
using System.Linq;
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
		public ActionResult Photos(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return SetDetails(model);
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