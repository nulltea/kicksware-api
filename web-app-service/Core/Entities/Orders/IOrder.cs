using Core.Attributes;

namespace Core.Entities
{
	[EntityService(Resource = "orders")]
	public interface IOrder : IBaseEntity { }
}