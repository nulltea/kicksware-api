using System.IO;
using Core.Reference;
using Newtonsoft.Json;

namespace web_app_service.Data.Reference_Data
{
	public static partial class Catalog
	{
		/// <summary>
		/// 
		/// </summary>
		public static readonly SneakerSize[] SneakerSizesList = JsonConvert.DeserializeObject<SneakerSize[]>(File.ReadAllText(@"Data\Json\sizes.json"));

		/// <summary>
		/// 
		/// </summary>
		public static readonly string[] SneakerBrandsList = JsonConvert.DeserializeObject<string[]>(File.ReadAllText(@"Data\Json\sneaker_brands.json"));
		
		/// <summary>
		/// 
		/// </summary>
		public static readonly string[] ColorsList = JsonConvert.DeserializeObject<string[]>(File.ReadAllText(@"Data\Json\colors.json"));
	}
}