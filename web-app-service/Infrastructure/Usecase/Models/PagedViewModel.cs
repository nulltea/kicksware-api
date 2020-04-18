using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities;
using Core.Exceptions;
using Core.Services;
using Microsoft.AspNetCore.Mvc;

namespace Infrastructure.Usecase.Models
{
	public class PagedModelList<T> : List<T> where T: IBaseEntity
	{
		private ICommonService<T> _service;

		[BindProperty(SupportsGet = true)]
		public int CurrentPage { get; set; }

		public int CountTotal { get; private set; }
		public int PagesTotal => (int)Math.Ceiling(decimal.Divide(CountTotal, PageSize));
		public int PageSize { get; }

		public bool HasPagePrevious => CurrentPage > 1;

		public bool HasPageNext => CurrentPage < PagesTotal;

		public PagedModelList(ICommonService<T> service, int currentPage = 1, int pageSize = 20)
		{
			_service = service;
			PageSize = pageSize;
			CurrentPage = currentPage;
			if (CurrentPage == 0) CurrentPage = 1;

			CountTotal = _service.Count();
		}

		public PagedModelList<T> FetchPage(int page)
		{
			if (0 >= page || page > PagesTotal) throw new PageNotValidException(page);
			CurrentPage = page;
			Clear();
			AddRange(_service.FetchOffset(PageSize, (page - 1) * PageSize));
			return this;
		}

		public PagedModelList<T> FetchNext()
		{
			if (!HasPageNext) throw new NextPageNotValidException();
			Clear();
			AddRange(_service.FetchOffset(PageSize, (++CurrentPage - 1) * PageSize));
			return this;
		}

		public PagedModelList<T> FetchPrevious()
		{
			if (!HasPagePrevious) throw new PreviousPageNotValidException();
			Clear();
			AddRange(_service.FetchOffset(PageSize, (--CurrentPage - 1) * PageSize));
			return this;
		}

		public async Task<PagedModelList<T>> FetchNextAsync()
		{
			if (!HasPageNext) throw new NextPageNotValidException();
			Clear();
			AddRange(await _service.FetchOffsetAsync(PageSize, (++CurrentPage - 1) * PageSize));
			return this;
		}

		public async Task<PagedModelList<T>> FetchPreviousAsync()
		{
			if (!HasPagePrevious) throw new PreviousPageNotValidException();
			Clear();
			AddRange(await _service.FetchOffsetAsync(PageSize, (--CurrentPage - 1) * PageSize));
			return this;
		}
	}

}