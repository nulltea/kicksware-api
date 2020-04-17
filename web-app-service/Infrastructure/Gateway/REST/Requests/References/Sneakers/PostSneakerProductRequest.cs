using System.Collections.Generic;
using Core.Entities.Products;
using Core.Entities.Reference;
using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class PostSneakerReferenceRequest : BaseSneakerReferenceRequest
	{
		public PostSneakerReferenceRequest(List<SneakerReference> sneakerReferences) : base(string.Empty, Method.POST)
		{
			AddJsonBody(sneakerReferences);
		}

		public PostSneakerReferenceRequest(SneakerReference sneakerReference) : base(string.Empty, Method.POST)
		{
			AddJsonBody(sneakerReference);
		}
	}
}