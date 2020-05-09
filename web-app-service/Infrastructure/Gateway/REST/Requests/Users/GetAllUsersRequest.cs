using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class GetAllUserRequest : UserBaseRequest
	{
		public GetAllUserRequest() : base(string.Empty) { }
	}
}