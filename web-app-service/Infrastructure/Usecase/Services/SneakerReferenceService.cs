using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Core.Entities;
using Core.Entities.References;
using Core.Extension;
using Core.Gateway;
using Core.Model.Parameters;
using Core.Reference;
using Core.Repositories;
using Core.Services;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Interact;
using Infrastructure.Pattern;

namespace Infrastructure.Usecase
{
	public class SneakerReferenceService : ISneakerReferenceService
	{
		private readonly ISneakerReferenceRepository _repository;

		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public SneakerReferenceService(ISneakerReferenceRepository repository, IGatewayClient<IGatewayRestRequest> client) =>
			(_repository, _client) = (repository, client);

		#region CRUD sync

		public SneakerReference FetchUnique(string sneakerId, RequestParams requestParams = default) =>
			_repository.GetUnique(sneakerId, requestParams);

		public List<SneakerReference> Fetch(RequestParams requestParams = default) => _repository.Get(requestParams);

		public List<SneakerReference> Fetch(IEnumerable<string> idList, RequestParams requestParams = default) => _repository.Get(idList, requestParams);

		public List<SneakerReference> Fetch(object queryObject, RequestParams requestParams = default) => _repository.Get(queryObject, requestParams);

		public List<SneakerReference> Fetch(Dictionary<string, object> queryMap, RequestParams requestParams = default) => _repository.Get(queryMap, requestParams);

		public SneakerReference Store(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.Post(sneakerReference, requestParams);

		public List<SneakerReference> Store(List<SneakerReference> sneakerReferences, RequestParams requestParams = default) => _repository.Post(sneakerReferences, requestParams);

		public bool Modify(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.Update(sneakerReference, requestParams);

		public int Count() => _repository.Count();

		public int Count(Dictionary<string, object> queryMap, RequestParams requestParams = default) => _repository.Count(queryMap, requestParams);

		public int Count(object queryObject, RequestParams requestParams = default) => _repository.Count(queryObject, requestParams);

		#endregion

		#region CRUD async

		public Task<SneakerReference> FetchUniqueAsync(string sneakerId, RequestParams requestParams = default) => _repository.GetUniqueAsync(sneakerId, requestParams);

		public Task<List<SneakerReference>> FetchAsync(RequestParams requestParams = default) => _repository.GetAsync(requestParams);

		public Task<List<SneakerReference>> FetchAsync(IEnumerable<string> idList, RequestParams requestParams = default) => _repository.GetAsync(idList, requestParams);

		public Task<List<SneakerReference>> FetchAsync(object queryObject, RequestParams requestParams = default) => _repository.GetAsync(queryObject, requestParams);

		public Task<List<SneakerReference>> FetchAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) => _repository.GetAsync(queryMap, requestParams);

		public Task<SneakerReference> StoreAsync(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.PostAsync(sneakerReference, requestParams);

		public Task<List<SneakerReference>> StoreAsync(List<SneakerReference> sneakerReferences, RequestParams requestParams = default) => _repository.PostAsync(sneakerReferences, requestParams);

		public Task<bool> ModifyAsync(SneakerReference sneakerReference, RequestParams requestParams = default) => _repository.UpdateAsync(sneakerReference, requestParams);

		public Task<int> CountAsync(Dictionary<string, object> queryMap, RequestParams requestParams = default) => _repository.CountAsync(queryMap, requestParams);

		public Task<int> CountAsync(object queryObject, RequestParams requestParams = default) => _repository.CountAsync(queryObject, requestParams);

		#endregion

		#region Usecases

		public List<SneakerReference> GetRelated(SneakerReference reference, RequestParams requestParams = default)
		{
			var requiredCount = requestParams?.Limit ?? 5;
			var query = new QueryBuilder()
				.SetQueryArguments("brandname", ExpressionType.And, new FilterParameter(reference.BrandName))
				.SetQueryArguments("modelname", ExpressionType.And, new FilterParameter(reference.ModelName, ExpressionType.Regex))
				.Build();

			if (requestParams?.Limit != null) requestParams.Limit++; // later -1 current reference

			var related = Fetch(query, requestParams);
			if ((related is null || related.Count < requiredCount) && reference.BaseModel != null)
			{
				query = new QueryBuilder()
					.SetQueryArguments("brandname", ExpressionType.And, new FilterParameter(reference.BrandName))
					.SetQueryArguments("modelname", ExpressionType.And, new FilterParameter(reference.BaseModel?.Name, ExpressionType.Regex))
					.Build();
				var lessRelated = Fetch(query);
				if (lessRelated != null && lessRelated.Any())
				{
					lessRelated = lessRelated.OrderBySimilarity(r => r.ModelName, reference.ModelName)
						.Where(r => r.UniqueID != reference.UniqueID)
						.ToList();
					related = (related?.Union(lessRelated) ?? lessRelated).ToList();
				}
			}

			if (related != null && related.Count < requiredCount && reference.Brand != null)
			{
				query = new QueryBuilder()
					.SetQueryArguments("brandname", ExpressionType.And, new FilterParameter(reference.BrandName))
					.Build();
				var lessRelated = Fetch(query);
				if (lessRelated != null && lessRelated.Any())
				{
					lessRelated = lessRelated.OrderBySimilarity(r => r.ModelName, reference.ModelName)
						.Where(r => r.UniqueID != reference.UniqueID)
						.ToList();
					related = related.Union(lessRelated).ToList();
				}
			}

			return related?
				.Where(r => r.UniqueID != reference.UniqueID)
				.Distinct(new EntityComparer<SneakerReference>())
				.Take(requiredCount).ToList();
		}

		public async Task<List<SneakerReference>> GetRelatedAsync(SneakerReference reference, RequestParams requestParams = default)
		{
			var requiredCount = requestParams?.Limit ?? 5;
			var query = new QueryBuilder()
				.SetQueryArguments("modelname", ExpressionType.And, new FilterParameter(reference.ModelName, ExpressionType.Regex))
				.Build();
			if (requestParams?.Limit != null) requestParams.Limit++; // later -1 current reference

			var related = await FetchAsync(query, requestParams);
			if ((related is null || related.Count < requiredCount) && reference.Model?.BaseModel != null)
			{
				query = new QueryBuilder()
					.SetQueryArguments("modelname", ExpressionType.And, new FilterParameter(reference.ModelName, ExpressionType.Regex))
					.Build();
				var lessRelated = await FetchAsync(query);
				if (lessRelated != null && lessRelated.Any())
				{
					lessRelated = lessRelated.OrderBySimilarity(r => r.ModelName, reference.ModelName)
						.Where(r => r.UniqueID != reference.UniqueID)
						.ToList();
					related = (related?.Union(lessRelated) ?? lessRelated).ToList();
				}
			}
			return related?
				.Where(r => r.UniqueID != reference.UniqueID)
				.Distinct(new EntityComparer<SneakerReference>())
				.Take(requiredCount).ToList();
		}

		public List<SneakerReference> GetFeatured(IEnumerable<string> models, RequestParams requestParams = default)
		{
			var query = new QueryBuilder()
				.SetQueryArguments("modelname", ExpressionType.Or, models.Select(model => new FilterParameter(model, ExpressionType.Regex)).ToArray())
				.Build();

			if (requestParams?.SortBy is null)
			{
				(requestParams ??= new RequestParams()).SortBy = "price";
				requestParams.SortDirection = SortDirection.Descending;
			}
			return Fetch(query, requestParams);
		}

		public async Task<List<SneakerReference>> GetFeaturedAsync(IEnumerable<string> models, RequestParams requestParams = default)
		{
			var query = new QueryBuilder()
				.SetQueryArguments("modelname", ExpressionType.And, models.Select(model => new FilterParameter(model, ExpressionType.Regex)).ToArray())
				.Build();

			if (requestParams?.SortBy is null)
			{
				(requestParams ??= new RequestParams()).SortBy = "price";
				requestParams.SortDirection = SortDirection.Descending;
			}
			return await FetchAsync(query, requestParams);
		}

		#endregion
	}
}