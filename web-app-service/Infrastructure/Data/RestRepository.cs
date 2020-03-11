using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities;
using Core.Repositories;
using Infrastructure.Communication.REST.Client;

namespace Infrastructure.Data
{

	//public class RestRepository<T> : IRepository<T>, IAsyncRepository<T> where T : IBaseEntity
	//{
	//	private readonly RestfulClient _client;

	//	protected RestRepository(RestfulClient client) => _client = client;


	//	public Task<T> GetUniqueAsync(string uniqueId)
	//	{

	//	}
		

	//	public Task<IReadOnlyList<T>> ListAllAsync()
	//	{

	//	}

	//	public Task<IReadOnlyList<T>> ListAsync(IEnumerable<string> ids)
	//	{

	//	}

	//	public Task<T> AddAsync(T entity)
	//	{

	//	}

	//	public Task UpdateAsync(T entity)
	//	{

	//	}

	//	public Task DeleteAsync(T entity)
	//	{

	//	}

	//	public Task<int> CountAsync()
	//	{

	//	}

	//	public T GetUnique(int id)
	//	{

	//	}

	//	public IReadOnlyList<T> ListAll()
	//	{

	//	}

	//	public IReadOnlyList<T> List()
	//	{

	//	}

	//	public T Post(T entity)
	//	{

	//	}

	//	public bool Update(T entity)
	//	{

	//	}

	//	public bool Delete(T entity)
	//	{

	//	}

	//	public int Count()
	//	{

	//	}
	//}

}