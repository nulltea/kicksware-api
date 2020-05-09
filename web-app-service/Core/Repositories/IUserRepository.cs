using Core.Entities.Users;

namespace Core.Repositories
{
	public interface IUserRepository : IAsyncRepository<User>, IRepository<User> { }
}