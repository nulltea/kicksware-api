using System.Collections.Generic;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class GetMapSneakersRequest : BaseSneakersListRequest
	{
		public GetMapSneakersRequest(Dictionary<string, object> map) : base("/map")
		{
			AddJsonBody(map);
		}

		public GetMapSneakersRequest(object map) : base("/map")
		{
			AddJsonBody(map);
		}
	}
}