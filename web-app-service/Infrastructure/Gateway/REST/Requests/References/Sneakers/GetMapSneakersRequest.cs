namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class GetMapSneakersRequest : BaseSneakersListRequest
	{
		public GetMapSneakersRequest(object map) : base("/map")
		{
			AddJsonBody(map);
		}
	}
}