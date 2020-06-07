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
	[EntityService(Resource = "products/sneakers")]
	public class SneakerProduct : IProduct
	{
		[Key]
		public string UniqueID { get; set; }

		public string BrandName { get; set; }

		public SneakerBrand Brand
		{
			get => _brand ??= BrandName;
			private set => _brand = value;
		}
		private SneakerBrand _brand;

		public string ModelName { get; set; }

		public SneakerModel Model
		{
			get => _model ??= new SneakerModel(ModelName, Brand);
			private set => _model = value;
		}
		private SneakerModel _model;

		public string ModelSKU { get; set; }

		public string ReferenceID { get; set; }

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

		public string Owner { get; set; }

		[JsonIgnore]
		public List<string> Photos { get; set; } = new List<string>();

		[JsonProperty("Images")]
		internal Dictionary<string, byte[]> Images = new Dictionary<string, byte[]>();

		public decimal ConditionIndex { get; set; }

		public Dictionary<string, ShippingInfo> ShippingInfo { get; set; }

		[JsonIgnore]
		public DateTime AddedAt { get; set; }

		public int Likes { get; set; }

		public bool Liked { get; set; }

		[OnDeserialized]
		internal void OnDeserialized(StreamingContext context)
		{
			foreach (var image in Images.Keys)
			{
				File.WriteAllBytes(Path.Combine(Constants.Constants.FileStoragePath, "products", image), Images[image]);
				Photos.Add(Path.Combine($"{Constants.Constants.FileStoragePath}/photos/products", image));
			}
		}

		public bool Equals(SneakerProduct other) => other != null && UniqueID == other.UniqueID;
	}
}