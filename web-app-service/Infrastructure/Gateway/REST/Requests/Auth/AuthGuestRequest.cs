namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthGuestRequest : AuthBaseRequest
	{
		public AuthGuestRequest() : base("/guest") { }
	}
}