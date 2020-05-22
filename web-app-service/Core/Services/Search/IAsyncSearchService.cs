using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Gateway;

namespace Core.Services
{
	public interface ISearchAsyncService<T>
	{
		Task<List<T>> SearchAsync(string query, RequestParams requestParams = default);

		Task<List<T>> SearchAsyncBy(string field, object query, RequestParams requestParams = default);
	}
}