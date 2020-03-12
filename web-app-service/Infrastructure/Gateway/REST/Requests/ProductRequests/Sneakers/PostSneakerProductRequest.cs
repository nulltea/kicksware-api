using Core.Entities.Products;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class PostSneakerProductRequest : BaseSneakerProductRequest 
	{
		public PostSneakerProductRequest(SneakerProduct sneakerProduct) : base(string.Empty, Method.POST)
		{
			AddJsonBody(sneakerProduct);
		}
	}
}