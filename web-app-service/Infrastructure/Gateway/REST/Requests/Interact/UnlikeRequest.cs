using RestSharp;

namespace Infrastructure.Gateway.REST.Interact
{
	public class UnlikeRequest : InteractBaseRequest
	{
		public UnlikeRequest(string entityID) : base("/unlike/{entityID}") =>
			AddParameter("entityID", entityID, ParameterType.UrlSegment);
	}
}