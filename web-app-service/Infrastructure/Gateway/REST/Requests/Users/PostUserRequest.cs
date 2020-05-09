using System.Collections.Generic;
using Core.Entities.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class PostUserRequest : UserBaseRequest
	{
		public PostUserRequest(User user) : base("/", Method.POST)
		{
			AddJsonBody(user);
		}

		public PostUserRequest(List<User> user) : base("/", Method.POST)
		{
			AddJsonBody(user);
		}
	}
}