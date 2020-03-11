namespace Infrastructure.Communication.REST.ProductRequests.Sneakers
{
	public class GetAllSneakersRequest : BaseSneakersListRequest
	{
		public GetAllSneakersRequest() : base("all") {}
	}
}