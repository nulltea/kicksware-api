using Core.Entities.Products;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class PutSneakerProductRequest : BaseSneakerProductRequest 
	{
		public PutSneakerProductRequest(SneakerProduct sneakerProduct) : base("/{sneakerId}", Method.PUT)
		{
			AddParameter("sneakerId", sneakerProduct.UniqueId, ParameterType.UrlSegment);
			AddJsonBody(sneakerProduct);
		}
	}
}