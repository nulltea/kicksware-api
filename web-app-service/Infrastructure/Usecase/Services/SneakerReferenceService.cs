using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Gateway;
using Core.Model.Parameters;
using Core.Reference;
using Core.Repositories;
using Core.Services;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;
using Infrastructure.Pattern;

namespace Infrastructure.Usecase
{
	public class SneakerReferenceService : ISneakerReferenceService
	{
		private readonly ISneakerReferenceRepository _repository;

		private readonly RestfulClient _client;

		public SneakerReferenceService(ISneakerReferenceRepository repository, RestfulClient client) => (_repository, _client) = (repository, client);

		#region CRUD Sync

		public SneakerReference FetchUnique(string sneakerId, RequestParams requestParams = default) =>
			_repository.GetUnique(sneakerId, requestParams);

		public List<SneakerReference> Fetch(RequestParams requestParams = default) => _repository.Get(requestParams);

		public List<SneakerReference> Fetch(IEnumerable<string> idList, RequestParams requestParams = default) => _repository.Get(idList, requestParams);

		public List<SneakerReference> Fetch(object queryObject, RequestParams requestParams = default) => _repository.Get(queryObject, requestParams);

		public List<SneakerReference> Fetch(Dictionary<string, object> queryMap, RequestParams requestParams = default) => _repository.Get(queryMap, requestParams);

		public SneakerReference Store(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.Post(sneakerReference, requestParams);

		public List<SneakerReference> Store(List<SneakerReference> sneakerReferences, RequestParams requestParams = default) => _repository.Post(sneakerReferences, requestParams);

		public bool Modify(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.Update(sneakerReference, requestParams);

		public int Count(Dictionary<string, object> queryMap = default, RequestParams requestParams = default) => _repository.Count(queryMap, requestParams);

		public int Count(object queryObject = default, RequestParams requestParams = default) => _repository.Count(queryObject, requestParams);

		#endregion

		#region CRUD Async

		public Task<SneakerReference> FetchUniqueAsync(string sneakerId, RequestParams requestParams = default) => _repository.GetUniqueAsync(sneakerId, requestParams);

		public Task<List<SneakerReference>> FetchAsync(RequestParams requestParams = default) => _repository.GetAsync(requestParams);

		public Task<List<SneakerReference>> FetchAsync(IEnumerable<string> idList, RequestParams requestParams = default) => _repository.GetAsync(idList, requestParams);

		public Task<List<SneakerReference>> FetchAsync(object queryObject, RequestParams requestParams = default) => _repository.GetAsync(queryObject, requestParams);

		public Task<List<SneakerReference>> FetchAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) => _repository.GetAsync(queryMap, requestParams);

		public Task<SneakerReference> StoreAsync(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.PostAsync(sneakerReference, requestParams);

		public Task<List<SneakerReference>> StoreAsync(List<SneakerReference> sneakerReferences, RequestParams requestParams = default) => _repository.PostAsync(sneakerReferences, requestParams);

		public Task<bool> ModifyAsync(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.UpdateAsync(sneakerReference, requestParams);

		public Task<int> CountAsync(Dictionary<string, object> queryMap = default, RequestParams requestParams = default) => _repository.CountAsync(queryMap, requestParams);

		public Task<int> CountAsync(object queryObject = default, RequestParams requestParams = default) => _repository.CountAsync(queryObject, requestParams);

		#endregion

		#region Usecases

		public List<SneakerReference> GetRelated(SneakerReference reference, RequestParams requestParams)
		{
			var query = new QueryBuilder()
				.SetQueryArguments("modelname", ExpressionType.And, new FilterParameter(reference.ModelName, ExpressionType.Regex))
				.Build();
			return Fetch(query, requestParams);
		}

		public Task<List<SneakerReference>> GetRelatedAsync(SneakerReference reference, RequestParams requestParams)
		{
			var query = new QueryBuilder()
				.SetQueryArguments("modelname", ExpressionType.And, new FilterParameter(reference.ModelName, ExpressionType.Regex))
				.Build();
			return FetchAsync(query, requestParams);
		}

		#endregion
	}
}