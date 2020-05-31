using Core.Entities.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthRemoteRequest : AuthBaseRequest
	{
		public AuthRemoteRequest(User user) : base("/remote", Method.POST)
		{
			AddJsonBody(user);
		}
	}
}