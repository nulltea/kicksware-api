using RestSharp;

namespace Infrastructure.Communication.REST.ProductRequests.Sneakers
{
	public class GetSneakerProductRequest : BaseSneakerProductRequest
	{
		public GetSneakerProductRequest(string sneakerId) : base("/{sneakerId}/")
		{
			AddParameter("sneakerId", sneakerId, ParameterType.UrlSegment);

		}
	}
}