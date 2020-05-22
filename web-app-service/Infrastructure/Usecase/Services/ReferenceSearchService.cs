using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.References;
using Core.Gateway;
using Core.Services;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Search.SneakerReference;

namespace Infrastructure.Usecase
{
	public class ReferenceSearchService : IReferenceSearchService
	{
		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public ReferenceSearchService(IGatewayClient<IGatewayRestRequest> client) => _client = client;

		#region Sync

		public List<SneakerReference> Search(string query, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new SearchReferenceRequest(query) {RequestParams = requestParams});

		public List<SneakerReference> SearchBy(string field, object query, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(
				new SearchReferenceByRequest(field, query) {RequestParams = requestParams});

		public List<SneakerReference> SearchSKU(string skuQuery, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new SearchReferenceBySKU(skuQuery) {RequestParams = requestParams});

		public List<SneakerReference> SearchBrand(string brandQuery, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(
				new SearchReferenceByBrandRequest(brandQuery) {RequestParams = requestParams});

		public List<SneakerReference> SearchModel(string modelQuery, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(
				new SearchReferenceByModelRequest(modelQuery) {RequestParams = requestParams});

		#endregion

		#region Async

		public Task<List<SneakerReference>> SearchAsync(string query, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(
				new SearchReferenceRequest(query) {RequestParams = requestParams});

		public Task<List<SneakerReference>>
			SearchAsyncBy(string field, object query, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(
				new SearchReferenceByRequest(field, query) {RequestParams = requestParams});

		public Task<List<SneakerReference>> SearchAsyncSKU(string skuQuery, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(
				new SearchReferenceBySKU(skuQuery) {RequestParams = requestParams});

		public Task<List<SneakerReference>>
			SearchAsyncBrand(string brandQuery, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(
				new SearchReferenceByBrandRequest(brandQuery) {RequestParams = requestParams});

		public Task<List<SneakerReference>>
			SearchAsyncModel(string modelQuery, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(
				new SearchReferenceByModelRequest(modelQuery) {RequestParams = requestParams});

		#endregion
	}
}