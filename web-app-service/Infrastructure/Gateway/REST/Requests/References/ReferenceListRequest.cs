using RestSharp;

namespace Infrastructure.Gateway.REST.References
{
	public class ReferenceListRequest : ReferenceRequest
	{
		public ReferenceListRequest(string referenceClass, string referenceQuery, Method method = Method.GET)
			: base(referenceClass, referenceQuery, method)
		{
			AddParameter("referenceClass", referenceClass, ParameterType.UrlSegment);
			AddParameter("referenceQuery", referenceQuery, ParameterType.QueryString);
		}
	}
}