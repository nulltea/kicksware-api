namespace Infrastructure.Gateway.REST.Mail
{
	public class PostPasswordResetRequest : MailBaseRequest
	{
		public PostPasswordResetRequest(string userID, string callbackUrl) : base("/password-reset")
		{
			AddParameter("userID", userID);
			AddParameter("callbackURL", callbackUrl);
		}
	}
}