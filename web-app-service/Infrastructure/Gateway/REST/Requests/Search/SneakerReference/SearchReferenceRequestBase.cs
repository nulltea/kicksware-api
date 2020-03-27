namespace Infrastructure.Gateway.REST.Search.SneakerReference
{
	public class SearchReferenceRequestBase : SearchRequest
	{
		public SearchReferenceRequestBase(string resource = default) : base("reference", resource) { }
	}
}