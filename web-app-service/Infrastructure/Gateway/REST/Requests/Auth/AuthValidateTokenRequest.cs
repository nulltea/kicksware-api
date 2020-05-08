using Core.Entities.Users;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthValidateTokenRequest : AuthBaseRequest
	{
		public AuthValidateTokenRequest(AuthToken token) : base("/token-validate") => AddParameter("token", token.Token);
	}
}