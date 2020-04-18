using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities;

namespace Core.Services
{
	public interface ICommonService<T> where T : IBaseEntity
	{
		T FetchOne(string uniqueId);

		List<T> FetchAll();

		List<T> FetchOffset(int count, int offset);

		int Count(object query = default);

		Task<T> FetchOneAsync(string uniqueId);

		Task<List<T>> FetchAllAsync();

		Task<List<T>> FetchOffsetAsync(int count, int offset);

		Task<int> CountAsync(object query = default);
	}
}