using System.Collections.Generic;
using System.Text.Json.Serialization;
using Core.Entities.Products;
using Microsoft.AspNetCore.Http;

namespace web_app_service.Models
{
	public class SneakerProductViewModel : SneakerProduct
	{
		[JsonIgnore]
		public List<IFormFile> FormFiles { get; set; }
	}
}
