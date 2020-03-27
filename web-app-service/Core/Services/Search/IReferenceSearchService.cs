using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Reference;

namespace Core.Services
{
	public interface IReferenceSearchService : ISearchService<SneakerReference>, ISearchAsyncService<SneakerReference>
	{
		#region Sync

		List<SneakerReference> SearchSKU(string skuQuery);

		List<SneakerReference> SearchBrand(string brandQuery);

		List<SneakerReference> SearchModel(string modelQuery);

		#endregion

		#region Async

		Task<List<SneakerReference>> SearchAsyncSKU(string skuQuery);

		Task<List<SneakerReference>> SearchAsyncBrand(string brandQuery);

		Task<List<SneakerReference>> SearchAsyncModel(string modelQuery);

		#endregion
	}
}