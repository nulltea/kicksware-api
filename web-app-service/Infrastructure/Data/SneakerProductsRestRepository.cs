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

		public SneakerProduct GetUnique(string sneakerId)
		{
			return _client.Request<SneakerProduct>(new GetSneakerProductRequest(sneakerId));
		}

		public List<SneakerProduct> GetAll()
		{
			return _client.Request<List<SneakerProduct>>(new GetAllSneakersRequest());
		}

		public List<SneakerProduct> Get(IEnumerable<string> idList)
		{
			return _client.Request<List<SneakerProduct>>(new GetQueriedSneakersRequest(idList));
		}

		public List<SneakerProduct> Get(object queryObject)
		{
			return _client.Request<List<SneakerProduct>>(new GetQueriedSneakersRequest(queryObject));
		}

		public SneakerProduct Post(SneakerProduct sneakerProduct)
		{
			return _client.Request<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct));
		}

		public bool Update(SneakerProduct sneakerProduct)
		{
			return _client.Request(new PutSneakerProductRequest(sneakerProduct));
		}

		public bool Delete(SneakerProduct sneakerProduct)
		{
			return _client.Request(new DeleteSneakerProductRequest(sneakerProduct));
		}

		public bool Delete(string sneakerId)
		{
			return _client.Request(new DeleteSneakerProductRequest(sneakerId));
		}

		public int Count(object queryObject)
		{
			return _client.Request<int>(new CountSneakerProductsRequest(queryObject));
		}

		#endregion

		#region Async

		public Task<SneakerProduct> GetUniqueAsync(string sneakerId)
		{
			return _client.RequestAsync<SneakerProduct>(new GetSneakerProductRequest(sneakerId));
		}

		public Task<List<SneakerProduct>> GetAllAsync()
		{
			return _client.RequestAsync<List<SneakerProduct>>(new GetAllSneakersRequest());
		}

		public Task<List<SneakerProduct>> GetAsync(IEnumerable<string> idList)
		{
			return _client.RequestAsync<List<SneakerProduct>>(new GetQueriedSneakersRequest(idList));
		}

		public Task<List<SneakerProduct>> GetAsync(object queryObject)
		{
			return _client.RequestAsync<List<SneakerProduct>>(new GetQueriedSneakersRequest(queryObject));
		}

		public Task<SneakerProduct> PostAsync(SneakerProduct sneakerProduct)
		{
			return _client.RequestAsync<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct));
		}

		public Task<bool> UpdateAsync(SneakerProduct sneakerProduct)
		{
			return _client.RequestAsync(new PostSneakerProductRequest(sneakerProduct));
		}

		public Task<bool> DeleteAsync(SneakerProduct sneakerProduct)
		{
			return _client.RequestAsync(new DeleteSneakerProductRequest(sneakerProduct));
		}

		public Task<bool> DeleteAsync(string sneakerId)
		{
			return _client.RequestAsync(new DeleteSneakerProductRequest(sneakerId));
		}

		public Task<int> CountAsync(object queryObject)
		{
			return _client.RequestAsync<int>(new CountSneakerProductsRequest(queryObject));
		}

		#endregion
	}
}