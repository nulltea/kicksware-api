using Core.Entities.Products;

namespace Core.Repositories
{
	public interface ISneakerProductRepository : IAsyncRepository<SneakerProduct>, IRepository<SneakerProduct>
	{ }
}