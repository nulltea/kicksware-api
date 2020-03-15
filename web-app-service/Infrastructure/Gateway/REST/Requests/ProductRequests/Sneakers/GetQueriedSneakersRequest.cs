using System.Collections;
using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class GetQueriedSneakersRequest : BaseSneakersListRequest
	{
		public GetQueriedSneakersRequest(IEnumerable<string> idCodes) : base("/query")
		{
			foreach (var id in idCodes)
			{
				AddParameter("sneakerId", id, ParameterType.QueryString);
			}
		}
	}
}