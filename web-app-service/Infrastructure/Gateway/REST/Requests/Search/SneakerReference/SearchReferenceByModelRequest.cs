using RestSharp;

namespace Infrastructure.Gateway.REST.Search.SneakerReference
{
	public class SearchReferenceByModelRequest : SearchReferenceRequestBase
	{
		public SearchReferenceByModelRequest(string modelQuery)
			: base("model/{model}") => AddParameter("model", modelQuery, ParameterType.UrlSegment);
	}
}