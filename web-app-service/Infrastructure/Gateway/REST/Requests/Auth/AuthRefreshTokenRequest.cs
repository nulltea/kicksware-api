using Core.Entities.Users;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthRefreshTokenRequest : AuthBaseRequest
	{
		public AuthRefreshTokenRequest(AuthToken token) : base("/token-refresh") => AddParameter("token", token.Token);
	}
}