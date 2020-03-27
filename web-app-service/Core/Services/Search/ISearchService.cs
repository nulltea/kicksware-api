using System.Collections.Generic;

namespace Core.Services
{
	public interface ISearchService<T>
	{
		List<T> Search(string query);

		List<T> SearchBy(string field, object query);
	}
}