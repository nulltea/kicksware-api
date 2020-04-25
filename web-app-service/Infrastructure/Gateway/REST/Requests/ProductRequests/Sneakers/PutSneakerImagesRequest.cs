using System.IO;
using Core.Entities.Products;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class PutSneakerImagesRequest : BaseSneakerProductRequest
	{
		public PutSneakerImagesRequest(SneakerProduct sneakerProduct) : base("/{sneakerId}/images", Method.PUT)
		{
			AddParameter("sneakerId", sneakerProduct.UniqueID, ParameterType.UrlSegment);
			AlwaysMultipartFormData = true;
			foreach (var photo in sneakerProduct.Photos)
			{
				AddFile(Path.GetFileName(photo), photo);
			}
		}
	}
}
