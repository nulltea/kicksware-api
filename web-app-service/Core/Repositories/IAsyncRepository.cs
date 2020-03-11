using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities;

namespace Core.Repositories
{
	public interface IAsyncRepository<T> where T : IBaseEntity
	{
		Task<T> GetUniqueAsync(string uniqueId);

		Task<IReadOnlyList<T>> ListAllAsync();

		Task<IReadOnlyList<T>> ListAsync(IEnumerable<string> idList); //todo condition

		Task<T> AddAsync(T entity);

		Task UpdateAsync(T entity);

		Task DeleteAsync(T entity);

		Task<int> CountAsync(); //todo condition
	}
}