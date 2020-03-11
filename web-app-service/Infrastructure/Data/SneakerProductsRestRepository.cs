using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities.Products;
using Core.Repositories;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;

namespace Infrastructure.Data
{
	public class SneakerProductsRestRepository : ISneakerProductRepository
	{
		private readonly RestfulClient _client;

		protected SneakerProductsRestRepository(RestfulClient client) => _client = client;

		#region Sync

		public SneakerProduct GetUnique(string sneakerId)
		{
			return _client.Request<SneakerProduct>(new GetSneakerProductRequest(sneakerId));
		}

		public IReadOnlyList<SneakerProduct> ListAll()
		{
			return _client.Request<IReadOnlyList<SneakerProduct>>(new GetAllSneakersRequest());
		}

		public IReadOnlyList<SneakerProduct> List(IEnumerable<string> idList)
		{
			return _client.Request<IReadOnlyList<SneakerProduct>>(new GetQueriedSneakersRequest(idList));
		}

		public IReadOnlyList<SneakerProduct> List(object queryObject)
		{
			return _client.Request<IReadOnlyList<SneakerProduct>>(new GetQueriedSneakersRequest(queryObject));
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

		public Task<IReadOnlyList<SneakerProduct>> ListAllAsync()
		{
			return _client.RequestAsync<IReadOnlyList<SneakerProduct>>(new GetAllSneakersRequest());
		}

		public Task<IReadOnlyList<SneakerProduct>> ListAsync(IEnumerable<string> idList)
		{
			return _client.RequestAsync<IReadOnlyList<SneakerProduct>>(new GetQueriedSneakersRequest(idList));
		}

		public Task<IReadOnlyList<SneakerProduct>> ListAsync(object queryObject)
		{
			return _client.RequestAsync<IReadOnlyList<SneakerProduct>>(new GetQueriedSneakersRequest(queryObject));
		}

		public Task<SneakerProduct> AddAsync(SneakerProduct sneakerProduct)
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