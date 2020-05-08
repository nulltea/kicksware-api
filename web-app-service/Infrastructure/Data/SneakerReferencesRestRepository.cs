using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.References;
using Core.Gateway;
using Core.Repositories;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.References.Sneakers;

namespace Infrastructure.Data
{
	public class SneakerReferencesRestRepository : ISneakerReferenceRepository
	{
		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public SneakerReferencesRestRepository(IGatewayClient<IGatewayRestRequest> client) => _client = client;

		#region Sync

		public SneakerReference GetUnique(string sneakerId, RequestParams requestParams = default) =>
			_client.Request<SneakerReference>(new GetSneakerReferenceRequest(sneakerId){RequestParams = requestParams});

		public List<SneakerReference> Get(RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new GetAllSneakerReferencesRequest{RequestParams = requestParams});

		public List<SneakerReference> Get(IEnumerable<string> idCodes, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(idCodes){RequestParams = requestParams});

		public List<SneakerReference> Get(object queryObject, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(queryObject){RequestParams = requestParams});

		public List<SneakerReference> Get(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(queryMap){RequestParams = requestParams});

		public SneakerReference Post(SneakerReference sneakerReference, RequestParams requestParams = default) =>
			_client.Request<SneakerReference>(new PostSneakerReferenceRequest(sneakerReference){RequestParams = requestParams});

		public List<SneakerReference> Post(List<SneakerReference> sneakerReferences, RequestParams requestParams = default) =>
			_client.Request<List<SneakerReference>>(new PostSneakerReferenceRequest(sneakerReferences){RequestParams = requestParams});

		public bool Update(SneakerReference sneakerReference, RequestParams requestParams = default) =>
			_client.Request(new PatchSneakerReferenceRequest(sneakerReference){RequestParams = requestParams});

		public bool Delete(SneakerReference sneakerProduct, RequestParams requestParams = default) => throw new NotImplementedException();

		public bool Delete(string referenceId, RequestParams requestParams = default) => throw new NotImplementedException();

		public int Count(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.Request<int>(new CountSneakerReferencesRequest(queryMap){RequestParams = requestParams});

		public int Count(object queryObject, RequestParams requestParams = default) =>
			_client.Request<int>(new CountSneakerReferencesRequest(queryObject){RequestParams = requestParams});

		public int Count() => _client.Request<int>(new CountSneakerReferencesRequest());

		#endregion

		#region Async

		public Task<SneakerReference> GetUniqueAsync(string sneakerId, RequestParams requestParams = default) =>
			_client.RequestAsync<SneakerReference>(new GetSneakerReferenceRequest(sneakerId){RequestParams = requestParams});

		public Task<List<SneakerReference>> GetAsync(RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(new GetAllSneakerReferencesRequest{RequestParams = requestParams});

		public Task<List<SneakerReference>> GetAsync(IEnumerable<string> idList, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(idList){RequestParams = requestParams});

		public Task<List<SneakerReference>> GetAsync(object queryObject, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(queryObject){RequestParams = requestParams});

		public Task<List<SneakerReference>> GetAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(queryMap){RequestParams = requestParams});

		public Task<SneakerReference> PostAsync(SneakerReference sneakerReference, RequestParams requestParams = default) =>
			_client.RequestAsync<SneakerReference>(new PostSneakerReferenceRequest(sneakerReference){RequestParams = requestParams});

		public Task<List<SneakerReference>> PostAsync(List<SneakerReference> sneakerReferences, RequestParams requestParams = default) =>
			_client.RequestAsync<List<SneakerReference>>(new PostSneakerReferenceRequest(sneakerReferences){RequestParams = requestParams});

		public Task<bool> UpdateAsync(SneakerReference sneakerReference, RequestParams requestParams = default) =>
			_client.RequestAsync(new PatchSneakerReferenceRequest(sneakerReference){RequestParams = requestParams});

		public Task<bool> DeleteAsync(SneakerReference sneakerProduct, RequestParams requestParams = default) => throw new NotImplementedException();

		public Task<bool> DeleteAsync(string referenceId, RequestParams requestParams = default) => throw new NotImplementedException();

		public  Task<int> CountAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) =>
			_client.RequestAsync<int>(new CountSneakerReferencesRequest(queryMap){RequestParams = requestParams});

		public Task<int> CountAsync(object queryObject, RequestParams requestParams = default) =>
			_client.RequestAsync<int>(new CountSneakerReferencesRequest(queryObject){RequestParams = requestParams});

		public Task<int> CountAsync() => _client.RequestAsync<int>(new CountSneakerReferencesRequest());

		#endregion
	}
}