using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class BaseSneakerReferenceRequest : ReferenceRequest
	{
		public BaseSneakerReferenceRequest(string resource, Method method = Method.GET)
			: base("sneakers", resource, method) { }
	}
}