using Core.Attributes;

namespace Core.Entities.Products
{
	[EntityService(Resource = "products")]
	public interface IProduct : IBaseEntity { }
}