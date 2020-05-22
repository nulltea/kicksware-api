using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.References;
using Core.Gateway;

namespace Core.Services
{
	public interface IReferenceSearchService : ISearchService<SneakerReference>, ISearchAsyncService<SneakerReference>
	{
		#region Sync

		List<SneakerReference> SearchSKU(string skuQuery, RequestParams requestParams = default);

		List<SneakerReference> SearchBrand(string brandQuery, RequestParams requestParams = default);

		List<SneakerReference> SearchModel(string modelQuery, RequestParams requestParams = default);

		#endregion

		#region Async

		Task<List<SneakerReference>> SearchAsyncSKU(string skuQuery, RequestParams requestParams = default);

		Task<List<SneakerReference>> SearchAsyncBrand(string brandQuery, RequestParams requestParams = default);

		Task<List<SneakerReference>> SearchAsyncModel(string modelQuery, RequestParams requestParams = default);

		#endregion
	}
}