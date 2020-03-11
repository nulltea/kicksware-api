using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests
{
	public class ProductRequest : RestRequest, IGatewayRestRequest
	{
		public ProductRequest(string productClass, string resource, Method method = Method.GET)
			: base("products/{productClass}" + resource, method)
		{
			AddParameter("productClass", productClass, ParameterType.UrlSegment);
		}
	}
}