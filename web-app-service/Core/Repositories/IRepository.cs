using System.Collections.Generic;
using Core.Entities;
using Core.Gateway;

namespace Core.Repositories
{
	public interface IRepository<T> where T : IBaseEntity
	{
		T GetUnique(string id, RequestParams requestParams = default);

		List<T> Get(RequestParams requestParams = default);

		List<T> Get(IEnumerable<string> idList, RequestParams requestParams = default);

		List<T> Get(Dictionary<string, object> queryMap, RequestParams requestParams = default);

		List<T> Get(object queryObject, RequestParams requestParams = default);

		T Post(T entity, RequestParams requestParams = default);

		bool Update(T entity, RequestParams requestParams = default);

		bool Delete(T entity, RequestParams requestParams = default);

		bool Delete(string uniqueId, RequestParams requestParams = default);

		int Count(Dictionary<string, object> queryMap, RequestParams requestParams = default);

		int Count(object queryObject, RequestParams requestParams = default);

		int Count();
	}
}