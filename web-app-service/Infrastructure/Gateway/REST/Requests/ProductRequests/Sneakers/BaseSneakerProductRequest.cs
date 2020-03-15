using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class BaseSneakerProductRequest : ProductRequest
	{
		public BaseSneakerProductRequest(string resource, Method method = Method.GET)
			: base("sneakers", resource, method) { }
	}
}