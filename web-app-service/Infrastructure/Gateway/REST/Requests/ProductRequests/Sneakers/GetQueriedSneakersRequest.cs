using System.Collections;
using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class GetQueriedSneakersRequest : BaseSneakersListRequest
	{
		public GetQueriedSneakersRequest(IEnumerable<string> idCodes) : base("")
		{
			foreach (var id in idCodes)
			{
				AddParameter("sneakerId", id, ParameterType.QueryString);
			}
		}

		public GetQueriedSneakersRequest(Dictionary<string, object> map) : base("/query", Method.POST)
		{
			AddJsonBody(map);
		}

		public GetQueriedSneakersRequest(object map) : base("/query", Method.POST)
		{
			AddJsonBody(map);
		}
	}
}