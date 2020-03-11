using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using Core.Enitities.Users;
using Core.Reference;

namespace Core.Enitities.Products
{
	/// <summary>
	/// Sneaker product enitity
	/// </summary>
	public class SneakerProduct : IProduct
	{
		[Key]
		public string UniqueId { get; set; }

		public string BrandName { get; set; }

		public string ModelName { get; set; }

		[DataType(DataType.Currency)]
		public decimal Price { get; set; }

		public SneakerType Type { get; set; }

		public SneakerSize Size { get; set; }

		public string Color { get; set; }

		public SneakerCondition Condition;

		public string Description { get; set; }

		public User Owner { get; set; }

		public List<string> Images => _images ??= new List<string>();

		public decimal ConditionIndex { get; set; }

		public DateTime AddedAt { get; set; }

		private List<string> _images = new List<string>();

	}
}