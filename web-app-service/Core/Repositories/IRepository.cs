using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities;

namespace Core.Repositories
{
	public interface IRepository<T> where T : IBaseEntity
	{
		T GetUnique(string id);

		IReadOnlyList<T> ListAll();

		IReadOnlyList<T> List(IEnumerable<string> idList); //todo condition

		T Post(T entity);

		bool Update(T entity);

		bool Delete(T entity);

		int Count(); //todo condition
	}
}