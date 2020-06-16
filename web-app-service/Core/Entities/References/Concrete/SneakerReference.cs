using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.IO;
using System.Linq;
using System.Net;
using System.Runtime.Serialization;
using Core.Attributes;
using Core.Entities.Products;
using Core.Reference;

namespace Core.Entities.References
{
	[EntityService(Resource = "references/sneakers")]
	public class SneakerReference : IReference
	{
		[Key]
		public string UniqueID { get; set; }

		public string ManufactureSku { get; set; }

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

		public string BaseModelName { get; set; }

		public SneakerModel BaseModel
		{
			get => _model ??= new SneakerModel(BaseModelName, Brand);
			private set => _model = value;
		}

		[DataType(DataType.Currency)]
		public decimal Price { get; set; }

		public string Description { get; set; }

		public string Color { get; set; }

		public Gender Gender { get; set; }

		public string Nickname { get; set; }

		public string Designer { get; set; }

		[DataType(DataType.ImageUrl)]
		public string ImageLink { get; set; }

		public List<string> ImageLinks { get; set; }

		public List<string> Materials { get; set; }

		public List<string> Categories { get; set; }

		public string ImagePath {
			get
			{
				if (string.IsNullOrEmpty(ImageLink)) return string.Empty; // TODO no image available icon
				return Path.Combine(Constants.Constants.FileStoragePath, "references", ImageLink);
			}
		}

		public List<string> OtherImages => ImageLinks.Select(img => Path.Combine(Constants.Constants.FileStoragePath, "references", img)).ToList();

		public DateTime ReleaseDate { get; set; }

		[DataType(DataType.Url)]
		public string StadiumUrl { get; set; }

		public int Likes { get; set; }

		public bool Liked { get; set; }

		[OnDeserialized]
		internal void OnDeserialized(StreamingContext context) { }

		public bool Equals(SneakerReference other) => other != null && UniqueID == other.UniqueID;
	}
}
