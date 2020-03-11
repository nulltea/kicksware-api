using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities.Products;
using Core.Repositories;
using Infrastructure.Communication.REST.Client;
using Infrastructure.Communication.REST.ProductRequests.Sneakers;

namespace Infrastructure.Data
{

	public class SneakerProductsRestRepository : ISneakerProductRepository
	{
		private readonly RestfulClient _client;

		protected SneakerProductsRestRepository(RestfulClient client) => _client = client;


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

		public Task<SneakerProduct> AddAsync(SneakerProduct sneakerProduct)
		{
			return _client.RequestAsync<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct));
		}

		public Task UpdateAsync(SneakerProduct entity)
		{
			return null;
		}

		public Task DeleteAsync(SneakerProduct entity)
		{
			return null;
		}

		public Task<int> CountAsync()
		{
			return null;
		}

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

		public SneakerProduct Post(SneakerProduct sneakerProduct)
		{
			return _client.Request<SneakerProduct>(new PostSneakerProductRequest(sneakerProduct));
		}

		public bool Update(SneakerProduct entity)
		{
			return true;
		}

		public bool Delete(SneakerProduct entity)
		{
			return true;
		}

		public int Count()
		{
			return 0;
		}
	}
}