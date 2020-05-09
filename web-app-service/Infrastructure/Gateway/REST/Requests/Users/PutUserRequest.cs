using Core.Entities.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class PutUserRequest : UserBaseRequest
	{
		public PutUserRequest(User user) : base("/", Method.PUT)
		{
			AddJsonBody(user);
		}
	}
}