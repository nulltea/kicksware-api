namespace Core.Entities.References
{
	public interface IModel : IBaseEntity
	{
		string Name { get; set; }

		string Description { get; set; }
	}
}