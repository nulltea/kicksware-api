using System;
using Newtonsoft.Json;

namespace Core.Entities.Users
{
	public class AuthToken
	{
		public string Name { get; set; }

		public string Token { get; set; }

		public bool Success { get; set; }

		public DateTime? Expires { get; set; }

		[JsonIgnore]
		public DateTimeOffset? ExpiresUtc => Expires is null ? null : new DateTimeOffset?(Expires.Value);


		public bool AllowRefresh { get; set; } = true;

		public AuthToken(string token) => Token = token;

		public AuthToken() { }

		public static implicit operator AuthToken(string token) => new AuthToken(token);

		public static implicit operator string(AuthToken token) => token.Token;
	}
}