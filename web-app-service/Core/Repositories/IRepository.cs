using System.Collections.Generic;
using Core.Entities;

namespace Core.Repositories
{
	public interface IRepository<T> where T : IBaseEntity
	{
		T GetUnique(string id);

		List<T> GetAll();

		List<T> GetOffset(int count, int offset);

		List<T> Get(IEnumerable<string> idList);

		List<T> Get(object queryObject);

		T Post(T entity);

		bool Update(T entity);

		bool Delete(T entity);

		bool Delete(string uniqueId);

		int Count(object queryObject = default);
	}
}