namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class CountSneakerProductsRequest : BaseSneakersListRequest
	{
		public CountSneakerProductsRequest(object queryObject) : base("/count?")
		{
			AddObject(queryObject);
		}
	}
}