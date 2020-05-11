
using Core.Entities.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthSingUpRequest : AuthBaseRequest
	{
		public AuthSingUpRequest(User user) : base("/sing-up", Method.POST)
		{
			AddJsonBody(user);
		}
	}
}