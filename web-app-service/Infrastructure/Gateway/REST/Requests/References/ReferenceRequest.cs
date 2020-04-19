using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.References
{
	public class ReferenceRequest : RestRequest, IGatewayRestRequest
	{
		public RequestParams RequestParams { get; set; }

		public ReferenceRequest(string referenceClass, string resource, Method method = Method.GET)
			: base("references/{referenceClass}" + resource, method)
		{
			AddParameter("referenceClass", referenceClass, ParameterType.UrlSegment);
		}
	}
}