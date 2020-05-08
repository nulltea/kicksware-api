using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Gateway;
using Core.Repositories;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;

namespace Infrastructure.Data
{
	public class SneakerProductsRestRepository : ISneakerProductRepository
	{
		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public SneakerProductsRestRepository(IGatewayClient<IGatewayRestRequest> client) => _client = client;

		#region Sync

		public SneakerProduct GetUnique(string sneakerId, RequestParams requestParams = default) =>
			_client.Request<SneakerProduct>(new GetSneakerProductRequest(sneakerId) {RequestParams = requestParams});

		public List<SneakerProduct> Get(RequestParams requestParams = default) =>
			_client.Request<List<SneakerProduct>>(new GetAllSneakersRequest {RequestParams = requestParams});

		public List<SneakerProduct> Get(IEnumerable<string> idCodes, RequestParams requestParams = default) =>
			_client.Request<List<SneakerProduct>>(new GetQueriedSneakersRequest(idCodes) {RequestParams = requestParams});

		public List<SneakerProduct> Get(object queryObject, RequestParams requestParams = default) =>
			_client.Request<List<SneakerProduct>>(new GetQueriedSneakersRequest(queryObject) {RequestParams = requestParams});

		public List<SneakerProduct> Get(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.Request<List<SneakerProduct>>(new GetQueriedSneakersRequest(queryMap) {RequestParams = requestParams});


		public SneakerProduct Post(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_client.Request<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct) {RequestParams = requestParams});

		public bool Update(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_client.Request(new PutSneakerProductRequest(sneakerProduct) {RequestParams = requestParams});

		public bool Delete(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_client.Request(new DeleteSneakerProductRequest(sneakerProduct) {RequestParams = requestParams});

		public bool Delete(string sneakerId, RequestParams requestParams = default) =>
			_client.Request(new DeleteSneakerProductRequest(sneakerId) {RequestParams = requestParams});

		public int Count(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			Get(queryMap, requestParams).Count;// TODO _client.Request<int>(new CountSneakerProductsRequest(queryObject));

		public int Count(object queryObject, RequestParams requestParams = default) =>
			Get(requestParams).Count;// TODO _client.Request<int>(new CountSneakerProductsRequest(queryObject));

		public int Count() => _client.Request<int>(new CountSneakerProductsRequest());

		#endregion

		#region Async

		public Task<SneakerProduct> GetUniqueAsync(string sneakerId, RequestParams requestParams = default) =>
			_client.RequestAsync<SneakerProduct>(new GetSneakerProductRequest(sneakerId) {RequestParams = requestParams});

		public Task<List<SneakerProduct>> GetAsync(RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerProduct>>(new GetAllSneakersRequest {RequestParams = requestParams});

		public Task<List<SneakerProduct>> GetAsync(IEnumerable<string> idList, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerProduct>>(new GetQueriedSneakersRequest(idList) {RequestParams = requestParams});

		public Task<List<SneakerProduct>> GetAsync(object queryObject, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerProduct>>(new GetQueriedSneakersRequest(queryObject) {RequestParams = requestParams});

		public Task<List<SneakerProduct>> GetAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerProduct>>(new GetQueriedSneakersRequest(queryMap) {RequestParams = requestParams});

		public Task<SneakerProduct> PostAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_client.RequestAsync<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct) {RequestParams = requestParams});

		public Task<bool> UpdateAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_client.RequestAsync(new PutSneakerProductRequest(sneakerProduct) {RequestParams = requestParams});

		public Task<bool> DeleteAsync(SneakerProduct sneakerProduct, RequestParams requestParams = default) =>
			_client.RequestAsync(new DeleteSneakerProductRequest(sneakerProduct) {RequestParams = requestParams});

		public Task<bool> DeleteAsync(string sneakerId, RequestParams requestParams = default) =>
			_client.RequestAsync(new DeleteSneakerProductRequest(sneakerId) {RequestParams = requestParams});

		public Task<int> CountAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.RequestAsync<int>(new CountSneakerProductsRequest(queryMap));

		public Task<int> CountAsync(object queryObject, RequestParams requestParams = default) =>
			_client.RequestAsync<int>(new CountSneakerProductsRequest(queryObject));

		public Task<int> CountAsync() => _client.RequestAsync<int>(new CountSneakerProductsRequest());

		#endregion
	}
}