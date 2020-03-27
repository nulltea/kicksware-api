using RestSharp;

namespace Infrastructure.Gateway.REST.Search.SneakerReference
{
	public class SearchReferenceByRequest : SearchReferenceRequestBase
	{
		public SearchReferenceByRequest(string field, object query) : base("by/{field}")
		{
			AddParameter("field", field, ParameterType.UrlSegment);
			AddParameter("query", query);
		}
	}
}