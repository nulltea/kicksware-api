using Newtonsoft.Json;

namespace Core.Entities.Users
{
	public class AuthClaims
	{
		[JsonProperty("exp")]
		public long ExpiresAt { get; set; }

		[JsonProperty("iat")]
		public long IssuedAt { get; set; }

		[JsonProperty("iss")]
		public string Issuer { get; set; }

		[JsonProperty("sub")]
		public string Subject { get; set; }

		[JsonProperty("gst")]
		public bool Guest { get; set; }

		[JsonProperty("adm")]
		public bool Admin { get; set; }
	}
}