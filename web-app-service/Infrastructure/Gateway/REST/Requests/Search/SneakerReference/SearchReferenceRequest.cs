namespace Infrastructure.Gateway.REST.Search.SneakerReference
{
	public class SearchReferenceRequest : SearchReferenceRequestBase
	{
		public SearchReferenceRequest(string query) => AddParameter("query", query);
	}
}