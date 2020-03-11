using Core.Communication;
using RestSharp;

namespace Infrastructure.Communication.REST.ProductRequests
{
	public class ProductsListRequest : ProductRequest, ICommunicationRequest
	{
		public ProductsListRequest(string productClass, string productsQuery, Method method = Method.GET)
			: base(productClass, productsQuery, method)
		{
			AddParameter("productClass", productClass, ParameterType.UrlSegment);
			AddParameter("productsQuery", productsQuery, ParameterType.QueryString);
		}
	}
}