using Core.Communication;
using RestSharp;

namespace Infrastructure.Communication.REST.ProductRequests
{
	public class ProductRequest : RestRequest, ICommunicationRequest
	{
		public ProductRequest(string productClass, string resource, Method method = Method.GET)
			: base("products/{productClass}" + resource, method)
		{
			AddParameter("productClass", productClass, ParameterType.UrlSegment);
		}
	}
}