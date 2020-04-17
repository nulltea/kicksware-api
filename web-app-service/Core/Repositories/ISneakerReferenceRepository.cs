using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.Reference;

namespace Core.Repositories
{
	public interface ISneakerReferenceRepository : IAsyncRepository<SneakerReference>, IRepository<SneakerReference>
	{
		List<SneakerReference> Post(List<SneakerReference> entities);

		Task<List<SneakerReference>> PostAsync(List<SneakerReference> entities);
	}
}