using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.Interact
{
	public class InteractBaseRequest : RestRequest, IGatewayRestRequest
	{
		public RequestParams RequestParams { get; set; }

		public InteractBaseRequest(string resource, Method method = Method.GET)
			: base("interact" + resource, method) { }
	}
}