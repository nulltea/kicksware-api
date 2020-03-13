using Core.Entities.Products;

namespace Core.Repositories
{
	public interface IUserRepository : IAsyncRepository<SneakerProduct>, IRepository<SneakerProduct> { }
}