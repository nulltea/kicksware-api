using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;
using Core.Entities.Products;

namespace Core.Entities.Reference
{
	public class SneakerReference : IProduct
	{
		[Key]
		public string UniqueId { get; set; }

		public string ManufactureSku { get; set; }

		public string BrandName { get; set; }

		public string ModelName { get; set; }

		[DataType(DataType.Currency)]
		public decimal Price { get; set; }

		public string Description { get; set; }

		public string Color { get; set; }

		public Gender Gender { get; set; }

		public string Nickname { get; set; }

		[DataType(DataType.ImageUrl)]
		public string ImageLink { get; set; }

		[DataType(DataType.Url)]
		public string StadiumUrl { get; set; }

		[OnDeserialized]
		internal void OnDeserialized(StreamingContext context) { }
	}
}