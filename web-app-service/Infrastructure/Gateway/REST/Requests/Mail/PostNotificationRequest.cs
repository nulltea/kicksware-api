namespace Infrastructure.Gateway.REST.Mail
{
	public class PostNotificationRequest : MailBaseRequest
	{
		public PostNotificationRequest(string userID, string callbackUrl) : base("/notify")
		{
			AddParameter("userID", userID);
			AddParameter("callbackURL", callbackUrl);
		}
	}
}