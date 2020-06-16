using System;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Text.RegularExpressions;
using Core.Attributes;
using Core.Extension;

namespace Core.Entities.References
{
	[EntityService(Resource = "references/brands")]
	public class SneakerBrand : IBrand
	{
		[Key]
		public string UniqueID { get; }

		public string Name { get; set; }

		public string Description { get; set; }

		public string Logo { get; set; }

		public string LogoPath => $"{Constants.Constants.FileStoragePath}/logos/{Logo ?? $"{UniqueID.ToLower()}-logo.svg"}";

		public string Hero { get; set; }

		public string HeroPath => $"{Constants.Constants.FileStoragePath}/heroes/{Hero ?? $"{UniqueID.ToLower()}-hero.jpg"}";

		public decimal Relevance { get; set; }

		public static implicit operator SneakerBrand(string field) => new SneakerBrand(field);

		public static implicit operator string(SneakerBrand property) => property.Name;

		public SneakerBrand(string name)
		{
			Name = name;
			UniqueID = Convert.ToString(name)?.ToFormattedID().ToLower();
		}

		public override string ToString() => Name;

		public bool Equals(SneakerBrand other) => other != null && UniqueID == other.UniqueID;
	}
}