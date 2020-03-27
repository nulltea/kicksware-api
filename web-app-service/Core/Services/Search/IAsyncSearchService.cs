using System.Collections.Generic;
using System.Threading.Tasks;

namespace Core.Services
{
	public interface ISearchAsyncService<T>
	{
		Task<List<T>> SearchAsync(string query);

		Task<List<T>> SearchAsyncBy(string field, object query);
	}
}