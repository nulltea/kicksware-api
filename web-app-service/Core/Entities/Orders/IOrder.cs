using Core.Attributes;

namespace Core.Entities
{
	[EntityService(Resource = "api/orders")]
	public interface IOrder : IBaseEntity { }
}