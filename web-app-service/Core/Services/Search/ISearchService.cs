using System.Collections.Generic;
using Core.Gateway;

namespace Core.Services
{
	public interface ISearchService<T>
	{
		List<T> Search(string query, RequestParams requestParams = default);

		List<T> SearchBy(string field, object query, RequestParams requestParams = default);
	}
}