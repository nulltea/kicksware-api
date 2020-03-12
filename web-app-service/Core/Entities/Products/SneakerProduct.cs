using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using Core.Entities.Users;
using Core.Reference;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace Core.Entities.Products
{
	/// <summary>
	/// Sneaker product entities
	/// </summary>
	public class SneakerProduct : IProduct
	{
		[Key]
		public string UniqueId { get; set; }

		public string BrandName { get; set; }

		public string ModelName { get; set; }

		[DataType(DataType.Currency)]
		public decimal Price { get; set; }

		[JsonConverter(typeof(StringEnumConverter))]
		public SneakerType Type { get; set; }

		public SneakerSize Size { get; set; }

		public string Color { get; set; }

		[JsonConverter(typeof(StringEnumConverter))]
		public SneakerCondition Condition { get; set; }

		public string Description { get; set; }

		public User Owner { get; set; }

		public List<string> Images => _images ??= new List<string>();

		public decimal ConditionIndex { get; set; }

		[JsonIgnore]
		public DateTime AddedAt { get; set; }

		private List<string> _images = new List<string>();

	}
}