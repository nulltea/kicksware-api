using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Repositories;
using Core.Services;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;

namespace Infrastructure.Usecase
{
	public class SneakerProductService : ISneakerProductService
	{
		private readonly ISneakerProductRepository _repository;

		private readonly RestfulClient _client;

		public SneakerProductService(ISneakerProductRepository repository, RestfulClient client) => (_repository, _client) = (repository, client);

		#region CRUD Sync

		public SneakerProduct FetchOne(string sneakerId) => _repository.GetUnique(sneakerId);

		public List<SneakerProduct> FetchAll() => _repository.GetAll();

		public List<SneakerProduct> Fetch(IEnumerable<string> idList) => _repository.Get(idList);

		public List<SneakerProduct> Fetch(object queryObject) => _repository.Get(queryObject);

		public SneakerProduct Store(SneakerProduct sneakerProduct)
		{
			var response = _repository.Post(sneakerProduct);

			if (response == null) return null;
			sneakerProduct.UniqueId = response.UniqueId;

			return !_client.Request(new PutSneakerImagesRequest(sneakerProduct)) ? null : response;
		}

		public bool Modify(SneakerProduct sneakerProduct) => _repository.Update(sneakerProduct);

		public bool Replace(SneakerProduct sneakerProduct) => _repository.Update(sneakerProduct);

		public bool Remove(SneakerProduct sneakerProduct) => _repository.Delete(sneakerProduct);

		public bool Remove(string sneakerId) => _repository.Delete(sneakerId);

		public int Count(object queryObject) => _repository.Count(queryObject);

		#endregion

		#region CRUD Async

		public Task<SneakerProduct> FetchOneAsync(string sneakerId) => _repository.GetUniqueAsync(sneakerId);

		public Task<List<SneakerProduct>> FetchAllAsync() => _repository.GetAllAsync();

		public Task<List<SneakerProduct>> FetchAsync(IEnumerable<string> idList) => _repository.GetAsync(idList);

		public Task<List<SneakerProduct>> FetchAsync(object queryObject) => _repository.GetAsync(queryObject);

		public async Task<SneakerProduct> StoreAsync(SneakerProduct sneakerProduct)
		{
			sneakerProduct = await _repository.PostAsync(sneakerProduct);

			if (sneakerProduct == null) return null;

			return !await _client.RequestAsync(new PutSneakerImagesRequest(sneakerProduct)) ? null : sneakerProduct;
		}

		public Task<bool> ModifyAsync(SneakerProduct sneakerProduct) => _repository.UpdateAsync(sneakerProduct);

		public Task<bool> ReplaceAsync(SneakerProduct sneakerProduct) => _repository.UpdateAsync(sneakerProduct);

		public Task<bool> RemoveAsync(SneakerProduct sneakerProduct) => _repository.DeleteAsync(sneakerProduct);

		public Task<bool> RemoveAsync(string sneakerId) => _repository.DeleteAsync(sneakerId);

		public Task<int> CountAsync(object queryObject) => _repository.CountAsync(queryObject);

		#endregion

		#region Usecases

		public bool AttachImages(SneakerProduct sneaker) => _client.Request(new PutSneakerImagesRequest(sneaker));

		public Task<bool> AttachImagesAsync(SneakerProduct sneaker) => _client.RequestAsync(new PutSneakerImagesRequest(sneaker));

		public Task<decimal> RequestConditionAnalysis(SneakerProduct sneaker) => throw new NotImplementedException();

		public Task<SneakerProduct> RequestSneakerPrediction(List<string> images) => throw new NotImplementedException();

		#endregion
	}
}