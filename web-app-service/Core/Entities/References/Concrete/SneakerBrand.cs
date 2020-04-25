using System;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Text.RegularExpressions;
using Core.Attributes;

namespace Core.Entities.References
{
	[EntityService(Resource = "api/references/brands")]
	public class SneakerBrand : IBrand
	{
		[Key]
		public string UniqueID { get; }

		public string Name { get; set; }

		public string Description { get; set; }

		public string Logo { get; set; }

		public decimal Relevance { get; set; }

		public string HeroPath { get; set; }

		public static implicit operator SneakerBrand(string field) => new SneakerBrand(field);

		public static implicit operator string(SneakerBrand property) => property.Name;

		public SneakerBrand(string name)
		{
			Name = name;
			UniqueID = new Regex("[\\n\\t;,.\\s()\\/]").Replace(Convert.ToString(name), "_").ToLower();
			Logo = $"logos/{UniqueID}-logo.svg";
			HeroPath = $"/images/heroes/{UniqueID}-hero.jpg";
		}

		public override string ToString() => Name;

		public bool Equals(SneakerBrand other) => other != null && UniqueID == other.UniqueID;
	}
}