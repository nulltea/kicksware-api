using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthBaseRequest : RestRequest, IGatewayRestRequest
	{
		public RequestParams RequestParams { get; set; }

		public AuthBaseRequest(string resource, Method method = Method.GET)
			: base("auth" + resource, method) { }
	}
}