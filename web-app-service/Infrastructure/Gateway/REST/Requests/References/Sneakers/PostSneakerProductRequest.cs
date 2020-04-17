using System.Collections.Generic;
using Core.Entities.Products;
using Core.Entities.Reference;
using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class PostManySneakerReferenceRequest : BaseSneakerReferenceRequest
	{
		public PostManySneakerReferenceRequest(List<SneakerReference> sneakerReferences) : base(string.Empty, Method.POST)
		{
			AddJsonBody(sneakerReferences);
		}
	}
}