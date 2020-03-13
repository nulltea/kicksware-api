using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Repositories;
using Core.Services;

namespace Infrastructure.Usecase
{
	public class SneakerProductService : ISneakerProductService
	{
		private readonly ISneakerProductRepository _repository;

		public SneakerProductService(ISneakerProductRepository repository) => _repository = repository;

		#region CRUD Sync

		public SneakerProduct RetrieveOne(string sneakerId) => _repository.GetUnique(sneakerId);

		public List<SneakerProduct> RetrieveAll() => _repository.GetAll();

		public List<SneakerProduct> Retrieve(IEnumerable<string> idList) => _repository.Get(idList);

		public List<SneakerProduct> Retrieve(object queryObject) => _repository.Get(queryObject);

		public SneakerProduct Store(SneakerProduct sneakerProduct) => _repository.Post(sneakerProduct);

		public bool Modify(SneakerProduct sneakerProduct) => _repository.Update(sneakerProduct);

		public bool Remove(SneakerProduct sneakerProduct) => _repository.Delete(sneakerProduct);

		public bool Remove(string sneakerId) => _repository.Delete(sneakerId);

		public int Count(object queryObject) => _repository.Count(queryObject);

		#endregion

		#region CRUD Async

		public Task<SneakerProduct> RetrieveOneAsync(string sneakerId) => _repository.GetUniqueAsync(sneakerId);

		public Task<List<SneakerProduct>> RetrieveAllAsync() => _repository.GetAllAsync();

		public Task<List<SneakerProduct>> RetrieveAsync(IEnumerable<string> idList) => _repository.GetAsync(idList);

		public Task<List<SneakerProduct>> RetrieveAsync(object queryObject) => _repository.GetAsync(queryObject);

		public Task<SneakerProduct> StoreAsync(SneakerProduct sneakerProduct) => _repository.PostAsync(sneakerProduct);

		public Task<bool> ModifyAsync(SneakerProduct sneakerProduct) => _repository.UpdateAsync(sneakerProduct);

		public Task<bool> RemoveAsync(SneakerProduct sneakerProduct) => _repository.DeleteAsync(sneakerProduct);

		public Task<bool> RemoveAsync(string sneakerId) => _repository.DeleteAsync(sneakerId);

		public Task<int> CountAsync(object queryObject) => _repository.CountAsync(queryObject);

		#endregion

		#region Usecases

		public Task<decimal> RequestConditionAnalysis(SneakerProduct sneaker) => throw new NotImplementedException();

		public Task<SneakerProduct> RequestSneakerPrediction(List<string> images) => throw new NotImplementedException();

		#endregion
	}
}