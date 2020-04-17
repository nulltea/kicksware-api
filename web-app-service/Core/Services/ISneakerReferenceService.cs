using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.Reference;
using Core.Repositories;

namespace Core.Services
{
	public interface ISneakerReferenceService
	{
		#region CRUD Sync

		SneakerReference FetchOne(string uniqueId);

		List<SneakerReference> FetchAll();

		List<SneakerReference> Fetch(IEnumerable<string> idList);

		List<SneakerReference> Fetch(object queryObject);

		SneakerReference Store(SneakerReference sneakerReference);

		List<SneakerReference> Store(List<SneakerReference> sneakerReferences);

		bool Modify(SneakerReference sneakerReference);

		int Count(object queryObject);

		#endregion

		#region CRUD Async

		Task<SneakerReference> FetchOneAsync(string uniqueId);

		Task<List<SneakerReference>> FetchAllAsync();

		Task<List<SneakerReference>> FetchAsync(IEnumerable<string> idList);

		Task<List<SneakerReference>> FetchAsync(object queryObject);

		Task<List<SneakerReference>> StoreAsync(List<SneakerReference> sneakerReferences);

		Task<SneakerReference> StoreAsync(SneakerReference sneakerReference);

		Task<bool> ModifyAsync(SneakerReference sneakerReference);

		Task<int> CountAsync(object queryObject);

		#endregion

		#region Usecases

		#endregion
	}
}