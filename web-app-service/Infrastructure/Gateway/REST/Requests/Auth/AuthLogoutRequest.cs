using Core.Entities.Users;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthLogoutRequest : AuthBaseRequest
	{
		public AuthLogoutRequest(AuthToken token) : base("/logout") => AddParameter("token", token.Token);
	}
}