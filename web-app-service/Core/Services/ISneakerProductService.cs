using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Repositories;

namespace Core.Services
{
	public interface ISneakerProductService
	{
		#region CRUD Sync

		SneakerProduct RetrieveOne(string uniqueId);

		List<SneakerProduct> RetrieveAll();

		List<SneakerProduct> Retrieve(IEnumerable<string> idList);

		List<SneakerProduct> Retrieve(object queryObject);

		SneakerProduct Store(SneakerProduct sneakerProduct);

		bool Modify(SneakerProduct sneakerProduct);

		bool Remove(SneakerProduct sneakerProduct);

		bool Remove(string uniqueId);

		int Count(object queryObject); //todo condition

		#endregion

		#region CRUD Async

		Task<SneakerProduct> RetrieveOneAsync(string uniqueId);

		Task<List<SneakerProduct>> RetrieveAllAsync();

		Task<List<SneakerProduct>> RetrieveAsync(IEnumerable<string> idList);

		Task<List<SneakerProduct>> RetrieveAsync(object queryObject);

		Task<SneakerProduct> StoreAsync(SneakerProduct sneakerProduct);

		Task<bool> ModifyAsync(SneakerProduct sneakerProduct);

		Task<bool> RemoveAsync(SneakerProduct sneakerProduct);

		Task<bool> RemoveAsync(string uniqueId);

		Task<int> CountAsync(object queryObject);

		#endregion

		#region Usecases

		Task<decimal> RequestConditionAnalysis(SneakerProduct sneaker);

		Task<SneakerProduct> RequestSneakerPrediction(List<string> images);

		#endregion
	}
}