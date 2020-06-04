using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Constants;
using Core.Entities;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Extension;
using Core.Model;
using Core.Model.Parameters;
using Core.Services;
using Core.Services.Interactive;
using Infrastructure.Usecase.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;
using SmartBreadcrumbs.Nodes;
using Web.Utils.Extensions;

namespace Web.Controllers
{
	public partial class ShopController : Controller
	{
		[ViewData]
		public string HeroCoverPath { get; set; } = $"{Constants.FileStoragePath}/heroes/shop-hero.jpg";

		[ViewData]
		public string HeroBreadTitle { get; set; } = "Buy sneakers";

		[ViewData]
		public string HeroBreadSubTitle { get; set; } = "Select and buy whatever kicks you like";

		[ViewData]
		public string HeroLogoPath { get; set; }

		private IFilteredModel<TEntity> InitFilterHandler<TEntity>(object additionalParams = default) where TEntity : IBaseEntity
		{
			var service = HttpContext.RequestServices.GetService<ICommonService<TEntity>>();
			var handler = new FilteredModel<TEntity>(service);

			var contentBuilder = HttpContext.RequestServices.GetService<FilterContentBuilder<TEntity>>();
			if (additionalParams != default) contentBuilder.SetAdditionalParams(additionalParams);
			contentBuilder.ConfigureFilter(handler);
			contentBuilder.ConfigureSorting(handler);
			return handler;
		}

		private IFilteredModel<IBaseEntity> InitFilterHandler(string entity, object additionalParams = default) =>
			entity switch
			{
				"references" => InitFilterHandler<SneakerReference>(additionalParams),
				"products" => InitFilterHandler<SneakerProduct>(additionalParams),
				"brand" => InitFilterHandler<SneakerReference>(additionalParams),
				"model" => InitFilterHandler<SneakerReference>(additionalParams),
				_ => InitFilterHandler<SneakerReference>(additionalParams),
			};

		[Route("shop/{entity}/requestUpdate/{entityID?}")]
		public async Task<IActionResult> RequestUpdate(string entity, string entityID, List<FilterInput> filterInputs, int page = 1,
														string sortBy = default)
		{
			var additionalParams = new Dictionary<string, object>();
			if (!string.IsNullOrEmpty(entityID)) additionalParams.Add($"{entity}ID", entityID);
			var handler = InitFilterHandler(entity, additionalParams);
			if (filterInputs != null && filterInputs.Any())
			{
				handler.ApplyUserInputs(filterInputs);
			}

			if (!string.IsNullOrEmpty(sortBy)) handler.ChooseSortParameter(sortBy);

			await handler.FetchPageAsync(page);

			return Json(new
			{
				content = await this.RenderViewAsync($"Partial/{entity.ToTitleCase()}Partial", handler, true),
				length = handler.CountTotal,
				pageSize = handler.PageSize
			});
		}

		private void AddBreadcrumbNode(string fromAction, string title)
		{
			var baseNode = new MvcBreadcrumbNode("References", "Shop", "Shop");
			var currentNode = new MvcBreadcrumbNode(fromAction, "Shop", title) {Parent = baseNode};
			ViewData["BreadcrumbNode"] = currentNode;
		}


		[Authorize(Policy = "NotGuest")]
		public async Task<IActionResult> Like(string id)
		{
			await HttpContext.RequestServices.GetService<ILikeService>().LikeAsync(id);
			return Ok();
		}

		[Authorize(Policy = "NotGuest")]
		public async Task<IActionResult> UnlikeAsync(string id)
		{
			await HttpContext.RequestServices.GetService<ILikeService>().UnlikeAsync(id);
			return Ok();
		}
	}
}