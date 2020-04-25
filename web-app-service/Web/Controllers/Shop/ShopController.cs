using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Extension;
using Core.Model;
using Core.Model.Parameters;
using Core.Services;
using Infrastructure.Usecase.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;
using SmartBreadcrumbs.Nodes;
using Web.Utils.Extensions;

namespace Web.Controllers
{
	public partial class ShopController : Controller
	{
		[ViewData]
		public string HeroCoverPath { get; set; } = "/images/heroes/shop-hero.jpg";

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

		private IFilteredModel<IBaseEntity> InitFilterHandler(string entity) =>
			entity switch
			{
				"references" => InitFilterHandler<SneakerReference>(),
				"products" => InitFilterHandler<SneakerProduct>(),
				"brand" => InitFilterHandler<SneakerReference>(), //TODO custom builder
				_ => InitFilterHandler<SneakerReference>(),
			};

		[Route("shop/{entity}/requestUpdate")]
		public async Task<IActionResult> RequestUpdate(string entity, List<FilterInput> filterInputs, int page = 1,
														string sortBy = default)
		{
			var handler = InitFilterHandler(entity);
			if (filterInputs != null && filterInputs.Any())
			{
				handler.ApplyUserInputs(filterInputs);
			}

			if (!string.IsNullOrEmpty(sortBy)) handler.ChooseSortParameter(sortBy);

			handler.FetchPage(page);

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
	}
}