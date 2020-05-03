using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities;
using Core.Gateway;

namespace Core.Repositories
{
	public interface IAsyncRepository<T> where T : IBaseEntity
	{
		Task<T> GetUniqueAsync(string uniqueId, RequestParams requestParams = default);

		Task<List<T>> GetAsync(RequestParams requestParams = default);

		Task<List<T>> GetAsync(IEnumerable<string> idList, RequestParams requestParams = default);

		Task<List<T>> GetAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default);

		Task<List<T>> GetAsync(object queryObject, RequestParams requestParams = default);

		Task<T> PostAsync(T entity, RequestParams requestParams = default);

		Task<bool> UpdateAsync(T entity, RequestParams requestParams = default);

		Task<bool> DeleteAsync(T entity, RequestParams requestParams = default);

		Task<bool> DeleteAsync(string uniqueId, RequestParams requestParams = default);

		Task<int> CountAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default);

		Task<int> CountAsync(object queryObject, RequestParams requestParams = default);

		Task<int> CountAsync();
	}
}