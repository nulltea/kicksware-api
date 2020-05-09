using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.Users
{
	public class UserBaseRequest : RestRequest, IGatewayRestRequest
	{
		public RequestParams RequestParams { get; set; }

		public UserBaseRequest(string resource, Method method = Method.GET) : base("users" + resource, method) { }
	}
}