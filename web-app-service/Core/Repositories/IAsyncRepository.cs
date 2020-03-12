using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities;

namespace Core.Repositories
{
	public interface IAsyncRepository<T> where T : IBaseEntity
	{
		Task<T> GetUniqueAsync(string uniqueId);

		Task<List<T>> GetAllAsync();

		Task<List<T>> GetAsync(IEnumerable<string> idList);

		Task<List<T>> GetAsync(object queryObject);

		Task<T> PostAsync(T entity);

		Task<bool> UpdateAsync(T entity);

		Task<bool> DeleteAsync(T entity);

		Task<bool> DeleteAsync(string uniqueId);

		Task<int> CountAsync(object queryObject); //todo condition
	}
}