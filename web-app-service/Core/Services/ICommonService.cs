using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities;
using Core.Gateway;

namespace Core.Services
{
	public interface ICommonService<T> where T : IBaseEntity
	{
		T FetchUnique(string uniqueId, RequestParams requestParams = default);

		List<T> Fetch(RequestParams requestParams = default);

		List<T> Fetch(Dictionary<string, object> query, RequestParams requestParams = default);

		int Count(Dictionary<string, object> query, RequestParams requestParams = default);

		int Count(object query = default, RequestParams requestParams = default);

		Task<T> FetchUniqueAsync(string uniqueId, RequestParams requestParams = default);

		Task<List<T>> FetchAsync(RequestParams requestParams = default);

		Task<List<T>> FetchAsync(Dictionary<string, object> query, RequestParams requestParams = default);

		Task<int> CountAsync(Dictionary<string, object> query, RequestParams requestParams = default);

		Task<int> CountAsync(object query = default, RequestParams requestParams = default);
	}
}