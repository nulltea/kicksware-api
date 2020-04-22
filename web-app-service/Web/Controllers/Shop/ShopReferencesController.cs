using Core.Entities.References;
using Microsoft.AspNetCore.Mvc;
using SmartBreadcrumbs.Attributes;

namespace Web.Controllers
{
	public partial class ShopController
	{
		[HttpGet]
		[Route("shop/references")]
		[Breadcrumb("Shop", FromAction = "Index", FromController = typeof(HomeController))]
		public IActionResult References(int page = 1, string sortBy = default)
		{
			var references = InitFilterHandler<SneakerReference>();
			if (!string.IsNullOrEmpty(sortBy)) references.ChooseSortParameter(sortBy);
			references.FetchPage(page);
			return View(references);
		}
	}
}