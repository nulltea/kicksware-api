using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class BaseSneakerReferenceListRequest : BaseSneakerReferenceRequest
	{
		public BaseSneakerReferenceListRequest(string referencesQuery, Method method = Method.GET)
			: base(referencesQuery, method) {}
	}
}