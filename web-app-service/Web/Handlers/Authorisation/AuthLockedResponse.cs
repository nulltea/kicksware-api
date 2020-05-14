namespace Web.Handlers.Authorisation
{
	public class AuthLockedResponse
	{
		public const string AuthLockedKey = "AuthLocked";

		public bool Locked { get; set; } = true;

		public string RedirectTo { get; set; }

		public AuthLockedResponse(string redirectTo = default) => RedirectTo = redirectTo;
	}
}