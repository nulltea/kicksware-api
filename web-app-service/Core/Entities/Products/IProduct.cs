using Core.Attributes;

namespace Core.Entities.Products
{
	[EntityService(Resource = "api/products")]
	public interface IProduct : IBaseEntity { }
}