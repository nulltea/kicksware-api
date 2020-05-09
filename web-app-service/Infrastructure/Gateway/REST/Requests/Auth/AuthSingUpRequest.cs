
using Core.Entities.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthSingUpRequest : AuthBaseRequest
	{
		public AuthSingUpRequest(User user, AuthCredentials credentials) : base("/sing-up", Method.POST)
		{
			AddJsonBody(credentials);
			AddJsonBody(user);
		}
	}
}