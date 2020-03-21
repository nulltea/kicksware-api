using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Linq;
using System.Runtime.Serialization;
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