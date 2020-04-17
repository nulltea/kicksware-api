using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.Reference;
using Core.Repositories;
using Core.Services;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;

namespace Infrastructure.Usecase
{
	public class SneakerReferenceService : ISneakerReferenceService
	{
		private readonly ISneakerReferenceRepository _repository;

		private readonly RestfulClient _client;

		public SneakerReferenceService(ISneakerReferenceRepository repository, RestfulClient client) => (_repository, _client) = (repository, client);

		#region CRUD Sync

		public SneakerReference FetchOne(string sneakerId) => _repository.GetUnique(sneakerId);

		public List<SneakerReference> FetchAll() => _repository.GetAll();

		public List<SneakerReference> Fetch(IEnumerable<string> idList) => _repository.Get(idList);

		public List<SneakerReference> Fetch(object queryObject) => _repository.Get(queryObject);

		public SneakerReference Store(SneakerReference sneakerReference) => _repository.Post(sneakerReference);

		public List<SneakerReference> Store(List<SneakerReference> sneakerReferences) => _repository.Post(sneakerReferences);

		public bool Modify(SneakerReference sneakerReference) => _repository.Update(sneakerReference);

		public int Count(object queryObject) => _repository.Count(queryObject);

		#endregion

		#region CRUD Async

		public Task<SneakerReference> FetchOneAsync(string sneakerId) => _repository.GetUniqueAsync(sneakerId);

		public Task<List<SneakerReference>> FetchAllAsync() => _repository.GetAllAsync();

		public Task<List<SneakerReference>> FetchAsync(IEnumerable<string> idList) => _repository.GetAsync(idList);

		public Task<List<SneakerReference>> FetchAsync(object queryObject) => _repository.GetAsync(queryObject);

		public Task<SneakerReference> StoreAsync(SneakerReference sneakerReference) => _repository.PostAsync(sneakerReference);

		public Task<List<SneakerReference>> StoreAsync(List<SneakerReference> sneakerReferences) => _repository.PostAsync(sneakerReferences);

		public Task<bool> ModifyAsync(SneakerReference sneakerReference) => _repository.UpdateAsync(sneakerReference);

		public Task<int> CountAsync(object queryObject) => _repository.CountAsync(queryObject);

		#endregion

		#region Usecases

		#endregion
	}
}