using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.References;
using Core.Gateway;
using Core.Services;
using Microsoft.AspNetCore.Mvc;
using Web.Models;
using Web.Utils;
using Web.Utils.Extensions;

namespace Web.Controllers
{
	public class SearchController : Controller
	{
		[HttpGet]
		public async Task<JsonResult> Search([FromServices] IReferenceSearchService service, [FromQuery] string prefix)
		{
			var results = service.Search(prefix, new RequestParams{Limit = 12});
			return Json(new
			{
				Success = true,
				Content = await this.RenderViewAsync("Header/_SearchResultsPartial", FormViewModel(results), true)
			});
		}

		private List<SearchResultViewModel> FormViewModel(List<SneakerReference> references)
		{
			return references.CastExtend<SneakerReference, SearchResultViewModel>();
		}
	}
}