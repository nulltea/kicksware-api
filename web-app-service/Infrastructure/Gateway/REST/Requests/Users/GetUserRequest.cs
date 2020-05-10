using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class GetUserRequest : UserBaseRequest
	{
		public GetUserRequest(string userID) : base("/{userID}")
		{
			AddParameter("userID", userID, ParameterType.UrlSegment);
		}
	}
}