using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Gateway;

namespace Core.Repositories
{
	public interface ISneakerReferenceRepository : IAsyncRepository<SneakerReference>, IRepository<SneakerReference>
	{
		List<SneakerReference> Post(List<SneakerReference> entities, RequestParams requestParams = default);

		Task<List<SneakerReference>> PostAsync(List<SneakerReference> entities, RequestParams requestParams = default);
	}
}