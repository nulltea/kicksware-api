using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Repositories;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;

namespace Infrastructure.Data
{
	public class SneakerProductsRestRepository : ISneakerProductRepository
	{
		private readonly RestfulClient _client;

		public SneakerProductsRestRepository(RestfulClient client) => _client = client;

		#region Sync

		public SneakerProduct GetUnique(string sneakerId) => _client.Request<SneakerProduct>(new GetSneakerProductRequest(sneakerId));

		public List<SneakerProduct> GetAll() => _client.Request<List<SneakerProduct>>(new GetAllSneakersRequest());

		public List<SneakerProduct> Get(IEnumerable<string> idCodes) => _client.Request<List<SneakerProduct>>(new GetQueriedSneakersRequest(idCodes));

		public List<SneakerProduct> Get(object queryObject) => _client.Request<List<SneakerProduct>>(new GetMapSneakersRequest(queryObject));

		public SneakerProduct Post(SneakerProduct sneakerProduct) => _client.Request<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct));

		public bool Update(SneakerProduct sneakerProduct) => _client.Request(new PutSneakerProductRequest(sneakerProduct));

		public bool Delete(SneakerProduct sneakerProduct) => _client.Request(new DeleteSneakerProductRequest(sneakerProduct));

		public bool Delete(string sneakerId) => _client.Request(new DeleteSneakerProductRequest(sneakerId));

		public int Count(object queryObject) => _client.Request<int>(new CountSneakerProductsRequest(queryObject));

		#endregion

		#region Async

		public Task<SneakerProduct> GetUniqueAsync(string sneakerId) => _client.RequestAsync<SneakerProduct>(new GetSneakerProductRequest(sneakerId));

		public Task<List<SneakerProduct>> GetAllAsync() => _client.RequestAsync<List<SneakerProduct>>(new GetAllSneakersRequest());

		public Task<List<SneakerProduct>> GetAsync(IEnumerable<string> idList) => _client.RequestAsync<List<SneakerProduct>>(new GetQueriedSneakersRequest(idList));

		public Task<List<SneakerProduct>> GetAsync(object queryObject) => _client.RequestAsync<List<SneakerProduct>>(new GetMapSneakersRequest(queryObject));

		public Task<SneakerProduct> PostAsync(SneakerProduct sneakerProduct) => _client.RequestAsync<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct));

		public Task<bool> UpdateAsync(SneakerProduct sneakerProduct) => _client.RequestAsync(new PostSneakerProductRequest(sneakerProduct));

		public Task<bool> DeleteAsync(SneakerProduct sneakerProduct) => _client.RequestAsync(new DeleteSneakerProductRequest(sneakerProduct));

		public Task<bool> DeleteAsync(string sneakerId) => _client.RequestAsync(new DeleteSneakerProductRequest(sneakerId));

		public Task<int> CountAsync(object queryObject) => _client.RequestAsync<int>(new CountSneakerProductsRequest(queryObject));

		#endregion
	}
}