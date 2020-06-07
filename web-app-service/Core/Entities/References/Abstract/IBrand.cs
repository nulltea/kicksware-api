using Core.Attributes;

namespace Core.Entities.References
{
	[EntityService(Resource = "references/brands")]
	public interface IBrand : IBaseEntity
	{
		string Name { get; set; }

		string Description { get; set; }

		string Logo { get; set; }
	}
}