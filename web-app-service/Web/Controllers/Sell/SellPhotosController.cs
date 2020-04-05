using System.IO;
using System.Linq;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		public ActionResult Photos(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(1, model);
			if (model.FormFiles != null && model.FormFiles.Any())
			{
				foreach (var formFile in model.FormFiles)
				{
					if (formFile.Length <= 0) continue;
					var fileName = Path.ChangeExtension(Path.GetRandomFileName(), Path.GetExtension(formFile.FileName));
					var filePath = Path.Combine(_environment.WebRootPath, "files", fileName);

					using var stream = System.IO.File.Create(filePath);
					formFile.CopyTo(stream);
					model.Photos.Add(filePath);
				}
			}

			model.Price = 500m;

			return this.ViewStep(3, model);
		}
	}
}