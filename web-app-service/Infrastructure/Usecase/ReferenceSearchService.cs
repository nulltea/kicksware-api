using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Reference;
using Core.Services;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.Search.SneakerReference;

namespace Infrastructure.Usecase
{
	public class ReferenceSearchService : IReferenceSearchService
	{
		private readonly RestfulClient _client;

		public ReferenceSearchService(RestfulClient client) => _client = client;

		#region Sync

		public List<SneakerReference> Search(string query) =>_client.Request<List<SneakerReference>>(new SearchReferenceRequest(query));

		public List<SneakerReference> SearchBy(string field, object query) => _client.Request<List<SneakerReference>>(new SearchReferenceByRequest(field, query));

		public List<SneakerReference> SearchSKU(string skuQuery) => _client.Request<List<SneakerReference>>(new SearchReferenceBySKU(skuQuery));

		public List<SneakerReference> SearchBrand(string brandQuery) => _client.Request<List<SneakerReference>>(new SearchReferenceByBrandRequest(brandQuery));

		public List<SneakerReference> SearchModel(string modelQuery) => _client.Request<List<SneakerReference>>(new SearchReferenceByModelRequest(modelQuery));

		#endregion

		#region Async

		public Task<List<SneakerReference>> SearchAsync(string query) => _client.RequestAsync<List<SneakerReference>>(new SearchReferenceRequest(query));

		public Task<List<SneakerReference>> SearchAsyncBy(string field, object query) => _client.RequestAsync<List<SneakerReference>>(new SearchReferenceByRequest(field, query));

		public Task<List<SneakerReference>> SearchAsyncSKU(string skuQuery) => _client.RequestAsync<List<SneakerReference>>(new SearchReferenceBySKU(skuQuery));

		public Task<List<SneakerReference>> SearchAsyncBrand(string brandQuery) => _client.RequestAsync<List<SneakerReference>>(new SearchReferenceByBrandRequest(brandQuery));

		public Task<List<SneakerReference>> SearchAsyncModel(string modelQuery) => _client.RequestAsync<List<SneakerReference>>(new SearchReferenceByModelRequest(modelQuery));

		#endregion
	}
}