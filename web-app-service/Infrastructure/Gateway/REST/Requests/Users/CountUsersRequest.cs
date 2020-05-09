using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class CountUsersRequest : UserBaseRequest
	{
		public CountUsersRequest() : base("/count") { }


		public CountUsersRequest(Dictionary<string, object> map) : base("/count", Method.POST)
		{
			AddJsonBody(map);
		}


		public CountUsersRequest(object map) : base("/count", Method.POST)
		{
			AddJsonBody(map);
		}
	}
}