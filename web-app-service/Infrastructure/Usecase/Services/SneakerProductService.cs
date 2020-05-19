using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Gateway;
using Core.Repositories;
using Core.Services;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.Interact;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;

namespace Infrastructure.Usecase
{
	public class SneakerProductService : ISneakerProductService
	{
		private readonly ISneakerProductRepository _repository;

		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public SneakerProductService(ISneakerProductRepository repository, IGatewayClient<IGatewayRestRequest> client) =>
			(_repository, _client) = (repository, client);

		#region CRUD sync

		public SneakerProduct FetchUnique(string sneakerId, RequestParams requestParams = default) =>
			_repository.GetUnique(sneakerId, requestParams);

		public List<SneakerProduct> Fetch(RequestParams requestParams = default) => _repository.Get(requestParams);

		public List<SneakerProduct> Fetch(IEnumerable<string> idList, RequestParams requestParams = default) =>
			_repository.Get(idList, requestParams);

		public List<SneakerProduct> Fetch(object queryObject, RequestParams requestParams = default) =>
			_repository.Get(queryObject, requestParams);

		public List<SneakerProduct> Fetch(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_repository.Get(queryMap, requestParams);

		public SneakerProduct Store(SneakerProduct sneakerProduct, RequestParams requestParams = default)
		{
			var response = _repository.Post(sneakerProduct, requestParams);

			if (response == null) return null;
			sneakerProduct.UniqueID = response.UniqueID;

			return !_client.Request(new PutSneakerImagesRequest(sneakerProduct)) ? null : response;
		}

		public bool Modify(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_repository.Update(sneakerProduct, requestParams);

		public bool Replace(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_repository.Update(sneakerProduct, requestParams);

		public bool Remove(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_repository.Delete(sneakerProduct, requestParams);

		public bool Remove(string sneakerId, RequestParams requestParams = default) => _repository.Delete(sneakerId, requestParams);

		public int Count(Dictionary<string, object> queryMap = default, RequestParams requestParams = default) =>
			_repository.Count(queryMap, requestParams);

		public int Count(object queryObject = default, RequestParams requestParams = default) =>
			_repository.Count(queryObject, requestParams);

		#endregion

		#region CRUD async

		public Task<SneakerProduct> FetchUniqueAsync(string sneakerId, RequestParams requestParams = default) =>
			_repository.GetUniqueAsync(sneakerId, requestParams);

		public Task<List<SneakerProduct>> FetchAsync(RequestParams requestParams = default) => _repository.GetAsync(requestParams);

		public Task<List<SneakerProduct>>
			FetchAsync(IEnumerable<string> idList, RequestParams requestParams = default) =>
			_repository.GetAsync(idList, requestParams);

		public Task<List<SneakerProduct>> FetchAsync(object queryObject, RequestParams requestParams = default) =>
			_repository.GetAsync(queryObject, requestParams);

		public Task<List<SneakerProduct>> FetchAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_repository.GetAsync(queryMap, requestParams);

		public async Task<SneakerProduct> StoreAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default)
		{
			sneakerProduct = await _repository.PostAsync(sneakerProduct, requestParams);

			if (sneakerProduct == null) return null;

			return !await _client.RequestAsync(new PutSneakerImagesRequest(sneakerProduct)) ? null : sneakerProduct;
		}

		public Task<bool> ModifyAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_repository.UpdateAsync(sneakerProduct, requestParams);

		public Task<bool> ReplaceAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_repository.UpdateAsync(sneakerProduct, requestParams);

		public Task<bool> RemoveAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_repository.DeleteAsync(sneakerProduct, requestParams);

		public Task<bool> RemoveAsync(string sneakerId, RequestParams requestParams = default) =>
			_repository.DeleteAsync(sneakerId, requestParams);

		public Task<int> CountAsync(Dictionary<string, object> queryMap = default, RequestParams requestParams = default) =>
			_repository.CountAsync(queryMap, requestParams);

		public Task<int> CountAsync(object queryObject = default, RequestParams requestParams = default) =>
			_repository.CountAsync(queryObject, requestParams);

		#endregion

		#region Usecases

		public bool AttachImages(SneakerProduct sneaker) =>
			_client.Request(new PutSneakerImagesRequest(sneaker));

		public Task<bool> AttachImagesAsync(SneakerProduct sneaker) =>
			_client.RequestAsync(new PutSneakerImagesRequest(sneaker));

		public Task<decimal> RequestConditionAnalysis(SneakerProduct sneaker) => throw new NotImplementedException();

		public Task<SneakerProduct> RequestSneakerPrediction(List<string> images) => throw new NotImplementedException();

		#endregion
	}
}