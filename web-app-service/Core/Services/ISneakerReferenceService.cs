using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Repositories;

namespace Core.Services
{
	public interface ISneakerProductService
	{
		#region CRUD Sync

		SneakerProduct FetchOne(string uniqueId);

		List<SneakerProduct> FetchAll();

		List<SneakerProduct> Fetch(IEnumerable<string> idList);

		List<SneakerProduct> Fetch(object queryObject);

		SneakerProduct Store(SneakerProduct sneakerProduct);

		bool Modify(SneakerProduct sneakerProduct);

		bool Replace(SneakerProduct sneakerProduct);

		bool Remove(SneakerProduct sneakerProduct);

		bool Remove(string uniqueId);

		int Count(object queryObject);

		#endregion

		#region CRUD Async

		Task<SneakerProduct> FetchOneAsync(string uniqueId);

		Task<List<SneakerProduct>> FetchAllAsync();

		Task<List<SneakerProduct>> FetchAsync(IEnumerable<string> idList);

		Task<List<SneakerProduct>> FetchAsync(object queryObject);

		Task<SneakerProduct> StoreAsync(SneakerProduct sneakerProduct);

		Task<bool> ModifyAsync(SneakerProduct sneakerProduct);

		Task<bool> ReplaceAsync(SneakerProduct sneakerProduct);

		Task<bool> RemoveAsync(SneakerProduct sneakerProduct);

		Task<bool> RemoveAsync(string uniqueId);

		Task<int> CountAsync(object queryObject);

		#endregion

		#region Usecases

		bool AttachImages(SneakerProduct sneaker);

		Task<bool> AttachImagesAsync(SneakerProduct sneaker);

		Task<decimal> RequestConditionAnalysis(SneakerProduct sneaker);

		Task<SneakerProduct> RequestSneakerPrediction(List<string> images);

		#endregion
	}
}