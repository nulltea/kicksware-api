using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities;

namespace Core.Repositories
{
	public interface IAsyncRepository<T> where T : IBaseEntity
	{
		Task<T> GetUniqueAsync(string uniqueId);

		Task<IReadOnlyList<T>> ListAllAsync();

		Task<IReadOnlyList<T>> ListAsync(IEnumerable<string> idList);

		Task<IReadOnlyList<T>> ListAsync(object queryObject);

		Task<T> AddAsync(T entity);

		Task<bool> UpdateAsync(T entity);

		Task<bool> DeleteAsync(T entity);

		Task<bool> DeleteAsync(string uniqueId);

		Task<int> CountAsync(object queryObject); //todo condition
	}
}