using RestSharp;

namespace Infrastructure.Gateway.REST.Interact
{
	public class LikeRequest : InteractBaseRequest
	{
		public LikeRequest(string entityID) : base("/like/{entityID}") =>
			AddParameter("entityID", entityID, ParameterType.UrlSegment);
	}
}