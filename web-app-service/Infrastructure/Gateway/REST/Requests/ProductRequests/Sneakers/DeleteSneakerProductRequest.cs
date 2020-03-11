using Core.Enitities.Products;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class DeleteSneakerProductRequest : BaseSneakerProductRequest 
	{
		public DeleteSneakerProductRequest(SneakerProduct sneakerProduct) : this(sneakerProduct.UniqueId) { }

		public DeleteSneakerProductRequest(string sneakerId) : base("/{sneakerId}", Method.DELETE)
		{
			AddParameter("sneakerId", sneakerId, ParameterType.UrlSegment);
		}
	}
}