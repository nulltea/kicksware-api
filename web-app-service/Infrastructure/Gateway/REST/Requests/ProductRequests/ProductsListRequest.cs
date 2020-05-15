using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests
{
	public abstract class ProductsListRequest : ProductRequest
	{
		public ProductsListRequest(string productClass, string productsQuery, Method method = Method.GET)
			: base(productClass, productsQuery, method)
		{
			AddParameter("productClass", productClass, ParameterType.UrlSegment);
			AddParameter("productsQuery", productsQuery, ParameterType.QueryString);
		}
	}
}