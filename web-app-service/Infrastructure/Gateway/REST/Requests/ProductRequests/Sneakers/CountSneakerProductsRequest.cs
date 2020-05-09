using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class CountSneakerProductsRequest : BaseSneakersListRequest
	{
		public CountSneakerProductsRequest() : base("/count") { }

		public CountSneakerProductsRequest(Dictionary<string, object> query) : base("/count", Method.POST)
		{
			AddJsonBody(query);
		}
		public CountSneakerProductsRequest(object query) : base("/count", Method.POST)
		{
			AddJsonBody(query);
		}
	}
}