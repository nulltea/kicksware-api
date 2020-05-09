using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class GetUserRequest : UserBaseRequest
	{
		public GetUserRequest(string username) : base("/{username}")
		{
			AddParameter("username", username, ParameterType.UrlSegment);
		}
	}
}