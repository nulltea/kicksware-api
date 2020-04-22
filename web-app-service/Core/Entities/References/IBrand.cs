using Core.Attributes;

namespace Core.Entities.References
{
	[EntityService(Resource = "api/references")]
	public interface IBrand : IBaseEntity
	{
		string Name { get; set; }

		string Description { get; set; }

		string Logo { get; set; }
	}
}