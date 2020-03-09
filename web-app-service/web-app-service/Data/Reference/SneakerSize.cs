using System.IO;
using Newtonsoft.Json;

namespace web_app_service.Data.Reference_Data
{
	public struct SneakerSize
	{
		public decimal Europe { get; set; }
		public decimal UnitedStates { get; set; }
		public decimal UnitedKingdom { get; set; }
		public decimal Centimeters { get; set; }
		public override string ToString() => $"{Europe} EU | {UnitedStates} US | {UnitedKingdom} UK";
	}

	public static partial class Catalog
	{
		/// <summary>
		/// 
		/// </summary>
		public static readonly SneakerSize[] SneakerSizesList = JsonConvert.DeserializeObject<SneakerSize[]>(File.ReadAllText(@"Data\Json\sizes.json"));
	}
}
 