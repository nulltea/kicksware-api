using RestSharp;

namespace Infrastructure.Gateway.REST.Search.SneakerReference
{
	public class SearchReferenceBySKU : SearchReferenceRequestBase
	{
		public SearchReferenceBySKU(string skuQuery)
			: base("sku/{sku}") => AddParameter("sku", skuQuery, ParameterType.UrlSegment);
	}
}