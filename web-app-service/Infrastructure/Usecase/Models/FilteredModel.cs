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

namespace Infrastructure.Usecase.Models
{
	public class FilteredModel<T> : List<T>, IFilteredModel<T> where T : IBaseEntity
	{
		private readonly ICommonService<T> _service;

		[BindProperty(SupportsGet = true)]
		public int CurrentPage { get; set; }

		public int CountTotal { get; private set; }
		public int PagesTotal => (int)Math.Ceiling(decimal.Divide(CountTotal, PageSize));
		public int PageSize { get; }

		public bool HasPagePrevious => CurrentPage > 1;

		public bool HasPageNext => CurrentPage < PagesTotal;

		public List<FilterGroup> FilterGroups { get; set; }

		public List<FilterParameter> FilterParameters => FilterGroups.SelectMany(g => g).ToList();

		public SortParameter SortParameter { get; set; }

		public FilteredModel(ICommonService<T> service, int currentPage = 1, int pageSize = 20)
		{
			_service = service;
			PageSize = pageSize;
			CurrentPage = currentPage;
			FilterGroups = new List<FilterGroup>();

			if (CurrentPage == 0) CurrentPage = 1;
		}

		public FilterGroup AddFilterGroup(string name, string property, ExpressionType expressionType = ExpressionType.In, string description = default)
		{
			var group = new FilterGroup(name, property, expressionType, description);
			FilterGroups.Add(group);
			return group;
		}

		public FilterGroup GetFilterGroup(string name) => FilterGroups.FirstOrDefault(g => g.Caption.ToLower().Equals(name.ToLower()));

		public FilterGroup this[string groupName]
		{
			get => GetFilterGroup(groupName);
			set => AddFilterGroup(value);
		}

		internal FilterGroup AddFilterGroup(FilterGroup group)
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

		public void ApplyUserInputs(Dictionary<string, (bool Checked, object Value)> filterInputs)
		{
			foreach (var filterInput in filterInputs)
			{
				var param = FilterParameters.FirstOrDefault(param => param.RenderId == filterInput.Key);
				if (param is null) continue;
				param.Checked = filterInput.Value.Checked;
				param.Value = filterInput.Value.Value;
			}
		}

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
				TakeCount = PageSize,
				SkipOffset = (page - 1) * PageSize
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
				TakeCount = PageSize,
				SkipOffset = (++CurrentPage - 1) * PageSize
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
			AddRange(_service.Fetch(new RequestParams
			{
				TakeCount = PageSize,
				SkipOffset = (--CurrentPage - 1) * PageSize
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
				TakeCount = PageSize,
				SkipOffset = (page - 1) * PageSize
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
			AddRange(await _service.FetchAsync(new RequestParams
			{
				TakeCount = PageSize,
				SkipOffset = (++CurrentPage - 1) * PageSize
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
			AddRange(await _service.FetchAsync(new RequestParams
			{
				TakeCount = PageSize,
				SkipOffset = (--CurrentPage - 1) * PageSize
			}));
		}

		private Dictionary<string, object> GetQueryMap()
		{
			var queryBuilder = new QueryRecourseBuilder();
			queryBuilder.SetQueryArguments(FilterGroups);
			return queryBuilder.Build();
		}

		private int GetCountTotal(Dictionary<string, object> query) => CountTotal = _service.Count(query);

		private async Task<int> GetCountTotalAsync(Dictionary<string, object> query) => CountTotal = await _service.CountAsync(query);
	}
}