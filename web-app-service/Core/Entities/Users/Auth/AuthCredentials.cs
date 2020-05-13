using System;

namespace Core.Entities.Users
{
	public class AuthCredentials
	{
		public string Email { get; set; }

		public string PasswordHash { get; set; }

		public AuthCredentials(string email, string password) => (Email, PasswordHash) = (email, password);

		public AuthCredentials() { }

		public static implicit operator AuthCredentials(ValueTuple<string, string> credentials) =>
			new AuthCredentials(credentials.Item1, credentials.Item2);

		public static implicit operator ValueTuple<string, string>(AuthCredentials credentials) =>
			(credentials.Email, credentials.PasswordHash);
	}
}