using System;

namespace Core.Entities.Users
{
	public class AuthCredentials
	{
		public string Username { get; set; }

		public string Password { get; set; }

		public AuthCredentials(string username, string password) => (Username, Password) = (username, password);

		public AuthCredentials() { }

		public static implicit operator AuthCredentials(ValueTuple<string, string> credentials) =>
			new AuthCredentials(credentials.Item1, credentials.Item2);

		public static implicit operator ValueTuple<string, string>(AuthCredentials credentials) =>
			(credentials.Username, credentials.Password);
	}
}