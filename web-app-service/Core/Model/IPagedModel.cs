using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities;

namespace Core.Model
{
	public interface IPagedModel<T> : IEnumerable<T> where T: IBaseEntity
	{
		int CurrentPage { get; set; }

		int CountTotal { get; }

		int PagesTotal { get; }

		int PageSize { get; }

		public bool HasPagePrevious { get; }

		public bool HasPageNext { get; }

		void FetchPage(int page);

		void FetchNext();

		void FetchPrevious();

		Task FetchPageAsync(int page);

		Task FetchNextAsync();

		Task FetchPreviousAsync();
	}
}