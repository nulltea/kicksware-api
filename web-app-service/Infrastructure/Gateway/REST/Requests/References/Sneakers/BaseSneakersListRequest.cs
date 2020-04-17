using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class BaseSneakersListRequest : BaseSneakerProductRequest
	{
		public BaseSneakersListRequest(string sneakersQuery, Method method = Method.GET)
			: base(sneakersQuery, method) {}
	}
}