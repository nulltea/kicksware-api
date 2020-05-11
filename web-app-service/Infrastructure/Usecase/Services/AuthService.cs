using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Gateway;
using Core.Services;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Auth;

namespace Infrastructure.Usecase
{
	public class AuthService : IAuthService
	{
		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public AuthService(IGatewayClient<IGatewayRestRequest> client) => _client = client;

		public bool SingUp(User user, AuthCredentials credentials, out AuthToken token) =>
			(token = _client.Request<AuthToken>(new AuthSingUpRequest(user))) != null;

		public bool Login(AuthCredentials credentials, out AuthToken token) =>
			(token = _client.Request<AuthToken>(new AuthLoginRequest(credentials))) != null;

		public bool Guest(out AuthToken token) =>
			(token = _client.Request<AuthToken>(new AuthGuestRequest())) != null;

		public void Logout(AuthToken token) => _client.Request(new AuthValidateTokenRequest(token));

		public bool RefreshToken(ref AuthToken token) =>
			(token = _client.Request<AuthToken>(new AuthRefreshTokenRequest(token))) != null;

		public bool ValidateToken(AuthToken token) => _client.Request<bool>(new AuthValidateTokenRequest(token));

		public Task<AuthToken> SingUpAsync(User user) =>
			_client.RequestAsync<AuthToken>(new AuthSingUpRequest(user));

		public Task<AuthToken> LoginAsync(AuthCredentials credentials) =>
			_client.RequestAsync<AuthToken>(new AuthLoginRequest(credentials));

		public Task<AuthToken> GuestAsync() =>
			 _client.RequestAsync<AuthToken>(new AuthGuestRequest());

		public Task LogoutAsync(AuthToken token) =>
			_client.RequestAsync(new AuthValidateTokenRequest(token));

		public Task<AuthToken> RefreshTokenAsync(AuthToken token) =>
			_client.RequestAsync<AuthToken>(new AuthRefreshTokenRequest(token));

		public Task<bool> ValidateTokenAsync(AuthToken token) =>
			_client.RequestAsync<bool>(new AuthValidateTokenRequest(token));
	}
}