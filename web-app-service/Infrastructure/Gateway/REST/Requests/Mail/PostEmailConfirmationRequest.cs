namespace Infrastructure.Gateway.REST.Mail
{
	public class PostEmailConfirmationRequest : MailBaseRequest
	{
		public PostEmailConfirmationRequest(string userID, string callbackUrl) : base("/confirm")
		{
			AddParameter("userID", userID);
			AddParameter("callbackURL", callbackUrl);
		}
	}
}