using System;
using Core.Attributes;

namespace Core.Entities.References
{
	[EntityService(Resource = "references")]
	public interface IReference : IBaseEntity
	{
		string ManufactureSku { get; set; }

		string BrandName { get; set; }

		string ModelName { get; set; }

		decimal Price { get; set; }

		string Description { get; set; }

		public DateTime ReleaseDate { get; set; }
	}
}