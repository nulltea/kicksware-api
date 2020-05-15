using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.Mail
{
	public abstract class MailBaseRequest : RestRequest, IGatewayRestRequest
	{
		public RequestParams RequestParams { get; set; }

		public MailBaseRequest(string resource, Method method = Method.GET)
			: base("mail" + resource, method) { }
	}
}