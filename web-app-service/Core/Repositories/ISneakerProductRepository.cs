using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Enitities.Products;

namespace Core.Repositories
{
	public interface ISneakerProductRepository : IAsyncRepository<SneakerProduct>, IRepository<SneakerProduct>
	{
		//decimal RequestConditionAnalysis(SneakerProduct sneaker);

		//SneakerProduct RequestSneakerPrediction(List<string> images);
	}
}