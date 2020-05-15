using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class GetSneakerReferenceRequest : BaseSneakerReferenceRequest
	{
		public GetSneakerReferenceRequest(string referenceId) : base("/{referenceId}")
		{
			AddParameter("referenceId", referenceId, ParameterType.UrlSegment);
		}
	}
}