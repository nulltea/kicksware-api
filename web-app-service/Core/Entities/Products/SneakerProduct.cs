using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Linq;
using System.Runtime.Serialization;
using Core.Attributes;
using Core.Entities.References;
using Core.Entities.Users;
using Core.Reference;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace Core.Entities.Products
{
	/// <summary>
	/// Sneaker product entities
	/// </summary>
	[EntityService(Resource = "api/products/sneakers")]
	public class SneakerProduct : IProduct
	{
		[Key]
		public string UniqueId { get; set; }

		public string BrandName { get; set; }

		public SneakerBrand Brand
		{
			get => _brand ??= BrandName;
			private set => _brand = value;
		}
		private SneakerBrand _brand;

		public string ModelName { get; set; }

		public string ModelSKU { get; set; }

		public string ModelRefId { get; set; }

		[DataType(DataType.Currency)]
		public decimal Price { get; set; }

		[DataType(DataType.Currency)]
		public decimal PriceOffset { get; set; }

		public bool AcceptOffers { get; set; } = false;

		[JsonConverter(typeof(StringEnumConverter))]
		public Currency Currency { get; set; } = Currency.UsDollar;

		[JsonConverter(typeof(StringEnumConverter))]
		public SneakerType Type { get; set; }

		public SneakerSize Size { get; set; }

		public string Color { get; set; }

		[JsonConverter(typeof(StringEnumConverter))]
		public SneakerCondition Condition { get; set; }

		public string Description { get; set; }

		public User Owner { get; set; }

		[JsonIgnore]
		public List<string> Photos { get; set; } = new List<string>();

		[JsonProperty("Images")]
		internal Dictionary<string, byte[]> Images = new Dictionary<string, byte[]>();

		public decimal ConditionIndex { get; set; }

		public Dictionary<string, ShippingInfo> ShippingInfo { get; set; }

		[JsonIgnore]
		public DateTime AddedAt { get; set; }

		[OnDeserialized]
		internal void OnDeserialized(StreamingContext context)
		{
			foreach (var image in Images.Keys.Reverse())
			{
				File.WriteAllBytes(Path.Combine(Constants.Constants.FileStoragePath, image), Images[image]);
				Photos.Add(image);
			}
		}
	}
}