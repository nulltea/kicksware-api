using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;
using Core.Entities.Products;
using Microsoft.AspNetCore.Http;

namespace Web.Models
{
	public class SneakerProductViewModel : SneakerProduct
	{
		[JsonIgnore]
		public List<IFormFile> FormFiles { get; set; }
	}
}
