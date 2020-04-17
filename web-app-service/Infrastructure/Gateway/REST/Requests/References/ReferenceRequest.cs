using RestSharp;

namespace Infrastructure.Gateway.REST.References
{
	public class ReferenceRequest : RestRequest, IGatewayRestRequest
	{
		public ReferenceRequest(string referenceClass, string resource, Method method = Method.GET)
			: base("products/{referenceClass}" + resource, method)
		{
			AddParameter("referenceClass", referenceClass, ParameterType.UrlSegment);
		}
	}
}