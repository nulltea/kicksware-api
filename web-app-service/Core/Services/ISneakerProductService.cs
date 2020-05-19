using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Gateway;
using Core.Repositories;
using Core.Services.Interactive;

namespace Core.Services
{
	public interface ISneakerProductService : ICommonService<SneakerProduct>
	{
		#region CRUD Sync

		List<SneakerProduct> Fetch(IEnumerable<string> idList, RequestParams requestParams = default);

		List<SneakerProduct> Fetch(object queryObject, RequestParams requestParams = default);

		SneakerProduct Store(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		bool Modify(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		bool Replace(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		bool Remove(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		bool Remove(string uniqueId, RequestParams requestParams = default);

		#endregion

		#region CRUD Async

		Task<List<SneakerProduct>> FetchAsync(IEnumerable<string> idList, RequestParams requestParams = default);

		Task<List<SneakerProduct>> FetchAsync(object queryObject, RequestParams requestParams = default);

		Task<SneakerProduct> StoreAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		Task<bool> ModifyAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		Task<bool> ReplaceAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		Task<bool> RemoveAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default);

		Task<bool> RemoveAsync(string uniqueId, RequestParams requestParams = default);

		#endregion

		#region Usecases

		bool AttachImages(SneakerProduct sneaker);

		Task<bool> AttachImagesAsync(SneakerProduct sneaker);

		Task<decimal> RequestConditionAnalysis(SneakerProduct sneaker);

		Task<SneakerProduct> RequestSneakerPrediction(List<string> images);

		#endregion
	}
}