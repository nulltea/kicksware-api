using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Gateway;
using Core.Repositories;
using Core.Services.Interactive;

namespace Core.Services
{
	public interface ISneakerReferenceService : ICommonService<SneakerReference>
	{
		#region CRUD Sync

		List<SneakerReference> Fetch(IEnumerable<string> idList, RequestParams requestParams = default);

		List<SneakerReference> Fetch(object queryObject, RequestParams requestParams = default);

		SneakerReference Store(SneakerReference sneakerReference, RequestParams requestParams = default);

		List<SneakerReference> Store(List<SneakerReference> sneakerReferences, RequestParams requestParams = default);

		bool Modify(SneakerReference sneakerReference, RequestParams requestParams = default);

		#endregion

		#region CRUD Async

		Task<List<SneakerReference>> FetchAsync(IEnumerable<string> idList, RequestParams requestParams = default);

		Task<List<SneakerReference>> FetchAsync(object queryObject, RequestParams requestParams = default);

		Task<List<SneakerReference>> StoreAsync(List<SneakerReference> sneakerReferences, RequestParams requestParams = default);

		Task<SneakerReference> StoreAsync(SneakerReference sneakerReference, RequestParams requestParams = default);

		Task<bool> ModifyAsync(SneakerReference sneakerReference, RequestParams requestParams = default);

		#endregion

		#region Usecases

		List<SneakerReference> GetRelated(SneakerReference reference, RequestParams requestParams = default);

		Task<List<SneakerReference>> GetRelatedAsync(SneakerReference reference, RequestParams requestParams = default);

		List<SneakerReference> GetFeatured(IEnumerable<string> models, RequestParams requestParams = default);

		Task<List<SneakerReference>> GetFeaturedAsync(IEnumerable<string> models, RequestParams requestParams = default);

		#endregion
	}
}