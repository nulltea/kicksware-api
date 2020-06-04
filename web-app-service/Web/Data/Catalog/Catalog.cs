using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Linq;
using System.Reflection;
using Core.Constants;
using Core.Reference;
using Newtonsoft.Json;

namespace Web.Data.Catalog
{
	public static partial class Catalog
	{
		/// <summary>
		///
		/// </summary>
		public static readonly SneakerSize[] SneakerSizesList = JsonConvert.DeserializeObject<SneakerSize[]>(File.ReadAllText(@$"{Constants.FileStoragePath}\meta\sizes.json"));

		/// <summary>
		///
		/// </summary>
		public static readonly string[] SneakerBrandsList = JsonConvert.DeserializeObject<string[]>(File.ReadAllText(@$"{Constants.FileStoragePath}\meta\sneaker_brands.json"));

		/// <summary>
		///
		/// </summary>
		public static readonly string[] ColorsList = JsonConvert.DeserializeObject<string[]>(File.ReadAllText(@$"{Constants.FileStoragePath}\meta\colors.json"));

		public static FilterColor[] FilterColors = JsonConvert.DeserializeObject<FilterColor[]>(File.ReadAllText(@$"{Constants.FileStoragePath}\meta\filter-colors.json"));

		public static Dictionary<string, string> CurrencySigns { get; } = Enum.GetValues(typeof(Currency)).OfType<Currency>().Select(value =>
		{
			var member = value.GetType().GetMember(value.ToString()).First();
			var code = Convert.ToString((int)value);
			var sign = member.GetCustomAttribute<DisplayAttribute>()?.ShortName;
			return (code, sign);
		}).Where(attr => !string.IsNullOrEmpty(attr.code) && !string.IsNullOrEmpty(attr.sign))
			.ToDictionary(attr => attr.code, attr => attr.sign);

		public static Dictionary<string, ShippingInfo> DefaultShippingInfo { get; } =
			new Dictionary<string, ShippingInfo>
			{
				{"United States",  new ShippingInfo(true, 30m)},
				{"United Kingdom", new ShippingInfo(true, 15m)},
				{"Europe",         new ShippingInfo(true, 15m)},
				{"Russia",         new ShippingInfo(true, 15m)},
				{"Canada",         new ShippingInfo(true, 30m)},
				{"Australia / NZ", new ShippingInfo(true, 30m)},
				{"Asia",           new ShippingInfo(true, 25m)},
				{"Other",          new ShippingInfo(true, 25m)},
			};
	}
}