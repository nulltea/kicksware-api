using System.Collections.Generic;
using Infrastructure.Gateway.REST.ProductRequests.Sneakers;
using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class GetQueriedUsersRequest : UserBaseRequest
	{
		public GetQueriedUsersRequest(IEnumerable<string> usernames) : base("/")
		{
			foreach (var username in usernames)
			{
				AddParameter("username", username, ParameterType.QueryString);
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