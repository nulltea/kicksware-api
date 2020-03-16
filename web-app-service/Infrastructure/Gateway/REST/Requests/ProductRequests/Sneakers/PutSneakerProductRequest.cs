using Core.Entities.Products;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class PutSneakerProductRequest : BaseSneakerProductRequest 
	{
		public PutSneakerProductRequest(SneakerProduct sneakerProduct) : base(string.Empty, Method.PUT)
		{
			AddJsonBody(sneakerProduct);
		}
	}
}