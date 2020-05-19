using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities;
using Core.Exceptions;
using Core.Gateway;
using Core.Model;
using Core.Model.Parameters;
using Core.Reference;
using Core.Services;
using Infrastructure.Pattern;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Infrastructure.Usecase.Models
{
	public class FilteredModel<T> : List<T>, IFilteredModel<T> where T : IBaseEntity
	{
		#region Properies

		[BindProperty(SupportsGet = true)]
		public int CurrentPage { get; set; }

		public int CountTotal { get; private set; }
		public int PagesTotal => Convert.ToInt32(Math.Ceiling(decimal.Divide(CountTotal, PageSize)));
		public int PageSize { get; }

		public bool HasPagePrevious => CurrentPage > 1;

		public bool HasPageNext => CurrentPage < PagesTotal;

		public List<FilterGroup> FilterGroups { get; }

		public List<FilterParameter> FilterParameters => FilterGroups.SelectMany(g => g).ToList();

		public List<SortParameter> SortParameters { get; set; }

		public SortParameter ChosenSorting { get; set; }

		private readonly ICommonService<T> _service;

		#endregion
		public FilteredModel(ICommonService<T> service, int currentPage = 1, int pageSize = 20)
		{
			_service = service;
			PageSize = pageSize;
			CurrentPage = currentPage;
			FilterGroups = new List<FilterGroup>();
			SortParameters = new List<SortParameter>();

			if (CurrentPage == 0) CurrentPage = 1;
		}

		#region Filtered

		public FilterGroup AddFilterGroup(string caption, FilterProperty property,
										ExpressionType expressionType = ExpressionType.In, string description = default)
		{
			var group = new FilterGroup(caption, property, expressionType, description);
			FilterGroups.Add(group);
			return group;
		}

		public FilterGroup AddForeignFilterGroup<TForeignEntity>(string caption, string fieldName,
																ExpressionType expressionType = ExpressionType.In,
																string description = default) =>
			AddForeignFilterGroup(caption, fieldName, typeof(TForeignEntity), expressionType, description);

		public FilterGroup AddForeignFilterGroup(string caption, string fieldName, Type foreignEntity,
												ExpressionType expressionType = ExpressionType.In,
												string description = default)
		{
			var group = new FilterGroup(caption, new FilterProperty(fieldName, foreignEntity), expressionType, description);
			FilterGroups.Add(group);
			return group;
		}

		public FilterGroup AddHiddenFilterGroup(FilterProperty property, ExpressionType expressionType = ExpressionType.In)
		{
			var group = new FilterGroup(property, expressionType) {Hidden = true};
			FilterGroups.Add(group);
			return group;
		}

		public FilterGroup AddHiddenFilterGroup<TForeignEntity>(string fieldName, ExpressionType expressionType = ExpressionType.In)
		{
			var group = new FilterGroup(new FilterProperty(fieldName, typeof(TForeignEntity)), expressionType)
			{
				Hidden = true
			};
			FilterGroups.Add(group);
			return group;
		}

		public FilterGroup GetFilterGroup(string name) => FilterGroups.FirstOrDefault(g => g.GroupID.Equals(name.ToLower()));

		public FilterGroup this[string groupName]
		{
			get => GetFilterGroup(groupName);
			set => AddFilterGroup(value);
		}

		private FilterGroup AddFilterGroup(FilterGroup group)
		{
			var existedGroup = GetFilterGroup(group.Caption);
			if (existedGroup is null)
			{
				FilterGroups.Add(group);
			}
			else
			{
				FilterGroups[FilterGroups.IndexOf(existedGroup)] = group;
			}
			return group;
		}

		public void ApplyUserInputs(List<FilterInput> filterInputs)
		{
			foreach (var input in filterInputs)
			{
				var param = FilterParameters.FirstOrDefault(p => p.RenderId == input.RenderId);
				if (param is null) continue;
				param.Checked = input.Checked;
				param.Value = input.Value;
			}
		}

		public void FetchFiltered() => FetchPage(1);

		private Dictionary<string, object> GetQueryMap()
		{
			var queryBuilder = new QueryBuilder();
			queryBuilder.SetQueryArguments(FilterGroups);
			return queryBuilder.Build();
		}

		#endregion

		#region Sorted

		public void AddSortParameters(params SortParameter[] parameters) => SortParameters.AddRange(parameters);

		public void AddSortParameters(Func<SortCriteria, SortParameter> selector)
		{
			SortParameters.AddRange(
				Enum.GetValues(typeof(SortCriteria)).Cast<SortCriteria>().Select(selector)
			);
		}

		public List<SelectListItem> FormSortSelect()
		{
			return SortParameters?.Select(sort =>
					new SelectListItem(sort.Caption, sort.RenderValue))
				.ToList() ?? new List<SelectListItem>();
		}

		public SortParameter ChooseSortParameter(string value) =>
			ChosenSorting = SortParameters.FirstOrDefault(sp => string.Equals(sp.RenderValue, value, StringComparison.CurrentCultureIgnoreCase));

		#endregion

		#region Paged

		public void FetchPage(int page)
		{
			var queryMap = GetQueryMap();
			if (GetCountTotal(queryMap) == 0)
			{
				Clear();
				return;
			}
			if (0 >= page || page > PagesTotal) throw new PageNotValidException(page);
			CurrentPage = page;
			Clear();
			AddRange(_service.Fetch(queryMap, new RequestParams
			{
				Limit = PageSize,
				Offset = (page - 1) * PageSize,
				SortBy = ChosenSorting?.Property.FieldName,
				SortDirection = ChosenSorting?.Direction
			}));
		}

		public void FetchNext()
		{
			var queryMap = GetQueryMap();
			if (GetCountTotal(queryMap) == 0)
			{
				Clear();
				return;
			}
			if (!HasPageNext) throw new NextPageNotValidException();
			Clear();
			AddRange(_service.Fetch(new RequestParams
			{
				Limit = PageSize,
				Offset = (++CurrentPage - 1) * PageSize,
				SortBy = ChosenSorting?.Property.FieldName,
				SortDirection = ChosenSorting?.Direction
			}));
		}

		public void FetchPrevious()
		{
			var queryMap = GetQueryMap();
			if (GetCountTotal(queryMap) == 0)
			{
				Clear();
				return;
			}
			if (!HasPagePrevious) throw new PreviousPageNotValidException();

			Clear();
			AddRange(_service.Fetch(queryMap, new RequestParams
			{
				Limit = PageSize,
				Offset = (--CurrentPage - 1) * PageSize,
				SortBy = ChosenSorting?.Property.FieldName,
				SortDirection = ChosenSorting?.Direction
			}));
		}

		public async Task FetchPageAsync(int page)
		{
			var queryMap = GetQueryMap();
			if (await GetCountTotalAsync(queryMap) == 0)
			{
				Clear();
				return;
			}
			if (0 >= page || page > PagesTotal) throw new PageNotValidException(page);
			CurrentPage = page;
			Clear();
			AddRange(await _service.FetchAsync(queryMap, new RequestParams
			{
				Limit = PageSize,
				Offset = (page - 1) * PageSize,
				SortBy = ChosenSorting?.Property.FieldName,
				SortDirection = ChosenSorting?.Direction
			}));
		}

		public async Task FetchNextAsync()
		{
			var queryMap = GetQueryMap();
			if (await GetCountTotalAsync(queryMap) == 0)
			{
				Clear();
				return;
			}
			if (!HasPageNext) throw new NextPageNotValidException();
			Clear();
			AddRange(await _service.FetchAsync(queryMap, new RequestParams
			{
				Limit = PageSize,
				Offset = (++CurrentPage - 1) * PageSize,
				SortBy = ChosenSorting?.Property.FieldName,
				SortDirection = ChosenSorting?.Direction
			}));
		}

		public async Task FetchPreviousAsync()
		{
			var queryMap = GetQueryMap();
			if (await GetCountTotalAsync(queryMap) == 0)
			{
				Clear();
				return;
			}
			if (!HasPagePrevious) throw new PreviousPageNotValidException();
			Clear();
			AddRange(await _service.FetchAsync(queryMap, new RequestParams
			{
				Limit = PageSize,
				Offset = (--CurrentPage - 1) * PageSize,
				SortBy = ChosenSorting?.Property.FieldName,
				SortDirection = ChosenSorting?.Direction
			}));
		}

		private int GetCountTotal(Dictionary<string, object> query) => CountTotal = _service.Count(query);

		private async Task<int> GetCountTotalAsync(Dictionary<string, object> query) => CountTotal = await _service.CountAsync(query);

		#endregion
	}
}