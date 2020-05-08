using Core.Entities.Users;

namespace Core.Services
{
	public interface IAuthService
	{
		bool SingUp(User user, AuthCredentials credentials, out AuthToken token);

		bool Login(AuthCredentials credentials, out AuthToken token);

		bool Guest(out AuthToken token);

		void Logout(AuthToken token);

		bool RefreshToken(ref AuthToken token);

		bool ValidateToken(AuthToken token);
	}
}