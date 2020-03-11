using System.Collections.Generic;
using Core.Enitities.Products;
using RestSharp;

namespace Infrastructure.Communication.REST.ProductRequests.Sneakers
{
	public class PostSneakerProductRequest : BaseSneakerProductRequest 
	{
		public PostSneakerProductRequest(SneakerProduct sneakerProduct) : base(string.Empty, Method.POST)
		{
			AddJsonBody(sneakerProduct);
		}
	}
}