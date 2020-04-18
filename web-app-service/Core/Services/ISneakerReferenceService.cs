using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.Reference;
using Core.Repositories;

namespace Core.Services
{
	public interface ISneakerReferenceService : ICommonService<SneakerReference>
	{
		#region CRUD Sync

		List<SneakerReference> Fetch(IEnumerable<string> idList);

		List<SneakerReference> Fetch(object queryObject);

		SneakerReference Store(SneakerReference sneakerReference);

		List<SneakerReference> Store(List<SneakerReference> sneakerReferences);

		bool Modify(SneakerReference sneakerReference);

		#endregion

		#region CRUD Async

		Task<List<SneakerReference>> FetchAsync(IEnumerable<string> idList);

		Task<List<SneakerReference>> FetchAsync(object queryObject);

		Task<List<SneakerReference>> StoreAsync(List<SneakerReference> sneakerReferences);

		Task<SneakerReference> StoreAsync(SneakerReference sneakerReference);

		Task<bool> ModifyAsync(SneakerReference sneakerReference);

		#endregion

		#region Usecases

		#endregion
	}
}