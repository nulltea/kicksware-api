using RestSharp;

namespace Infrastructure.Gateway.REST.Search.SneakerReference
{
	public class SearchReferenceByBrandRequest : SearchReferenceRequestBase
	{
		public SearchReferenceByBrandRequest(string brandQuery)
			: base("brand/{brand}") => AddParameter("brand", brandQuery, ParameterType.UrlSegment);
	}
}