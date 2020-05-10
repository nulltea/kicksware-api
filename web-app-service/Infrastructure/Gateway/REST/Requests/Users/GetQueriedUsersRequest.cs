using System.Collections.Generic;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;
using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class GetQueriedUsersRequest : UserBaseRequest
	{
		public GetQueriedUsersRequest(IEnumerable<string> userIDs) : base("/")
		{
			foreach (var userID in userIDs)
			{
				AddParameter("userID", userID, ParameterType.QueryString);
			}
		}

		public GetQueriedUsersRequest(Dictionary<string, object> map) : base("/query", Method.POST)
		{
			AddJsonBody(map);
		}


		public GetQueriedUsersRequest(object map) : base("/query", Method.POST)
		{
			AddJsonBody(map);
		}
	}
}