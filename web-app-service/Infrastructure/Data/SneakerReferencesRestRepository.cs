using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Reference;
using Core.Repositories;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;
using Infrastructure.Gateway.REST.References.Sneakers;

namespace Infrastructure.Data
{
	public class SneakerReferencesRestRepository : ISneakerReferenceRepository
	{
		private readonly RestfulClient _client;

		public SneakerReferencesRestRepository(RestfulClient client) => _client = client;

		#region Sync

		public SneakerReference GetUnique(string sneakerId) => _client.Request<SneakerReference>(new GetSneakerReferenceRequest(sneakerId));

		public List<SneakerReference> GetAll() => _client.Request<List<SneakerReference>>(new GetAllSneakerReferencesRequest());

		public List<SneakerReference> GetOffset(int count, int offset) => GetAll().Skip(offset).Take(count).ToList(); //TODO rest

		public List<SneakerReference> Get(IEnumerable<string> idCodes) => _client.Request<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(idCodes));

		public List<SneakerReference> Get(object queryObject) => _client.Request<List<SneakerReference>>(new GetMapSneakerReferencesRequest(queryObject));

		public SneakerReference Post(SneakerReference sneakerReference) => _client.Request<SneakerReference>(new PostSneakerReferenceRequest(sneakerReference));

		public List<SneakerReference> Post(List<SneakerReference> sneakerReferences) => _client.Request<List<SneakerReference>>(new PostSneakerReferenceRequest(sneakerReferences));

		public bool Update(SneakerReference sneakerReference) => _client.Request(new PatchSneakerReferenceRequest(sneakerReference));

		public bool Delete(SneakerReference sneakerProduct) => throw new NotImplementedException();

		public bool Delete(string referenceId) => throw new NotImplementedException();

		public int Count(object queryObject) => GetAll().Count;// TODO _client.Request<int>(new CountSneakerReferencesRequest(queryObject));

		#endregion

		#region Async

		public Task<SneakerReference> GetUniqueAsync(string sneakerId) => _client.RequestAsync<SneakerReference>(new GetSneakerReferenceRequest(sneakerId));

		public Task<List<SneakerReference>> GetAllAsync() => _client.RequestAsync<List<SneakerReference>>(new GetAllSneakerReferencesRequest());

		public async Task<List<SneakerReference>> GetOffsetAsync(int count, int offset) => (await GetAllAsync()).Skip(offset).Take(count).ToList(); //TODO rest

		public Task<List<SneakerReference>> GetAsync(IEnumerable<string> idList) => _client.RequestAsync<List<SneakerReference>>(new GetQueriedSneakerReferencesRequest(idList));

		public Task<List<SneakerReference>> GetAsync(object queryObject) => _client.RequestAsync<List<SneakerReference>>(new GetMapSneakerReferencesRequest(queryObject));

		public Task<SneakerReference> PostAsync(SneakerReference sneakerReference) => _client.RequestAsync<SneakerReference>(new PostSneakerReferenceRequest(sneakerReference));

		public Task<List<SneakerReference>> PostAsync(List<SneakerReference> sneakerReferences) => _client.RequestAsync<List<SneakerReference>>(new PostSneakerReferenceRequest(sneakerReferences));

		public Task<bool> UpdateAsync(SneakerReference sneakerReference) => _client.RequestAsync(new PatchSneakerReferenceRequest(sneakerReference));

		public Task<bool> DeleteAsync(SneakerReference sneakerProduct) => throw new NotImplementedException();

		public Task<bool> DeleteAsync(string referenceId) => throw new NotImplementedException();

		public async Task<int> CountAsync(object queryObject) => (await GetAllAsync()).Count;// TODO _client.RequestAsync<int>(new CountSneakerReferencesRequest(queryObject));

		#endregion
	}
}