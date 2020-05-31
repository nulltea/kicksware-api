using System.Threading.Tasks;
using Core.Entities.Users;

namespace Core.Services
{
	public interface IAuthService
	{
		bool SingUp(User user, AuthCredentials credentials, out AuthToken token);

		bool Login(AuthCredentials credentials, out AuthToken token);

		bool Remote(User user, out AuthToken token);

		bool Guest(out AuthToken token);

		void Logout(AuthToken token);

		bool RefreshToken(ref AuthToken token);

		bool ValidateToken(AuthToken token);

		Task<AuthToken> SingUpAsync(User user);

		Task<AuthToken> LoginAsync(AuthCredentials credentials);

		Task<AuthToken> RemoteAsync(User user);

		Task<AuthToken> GuestAsync();

		Task LogoutAsync(AuthToken token);

		Task<AuthToken> RefreshTokenAsync(AuthToken token);

		Task<bool> ValidateTokenAsync(AuthToken token);
	}
}