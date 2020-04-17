using System.IO;
using Core.Entities.Products;
using Core.Entities.Reference;
using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class PatchSneakerReferenceRequest : BaseSneakerReferenceRequest
	{
		public PatchSneakerReferenceRequest(SneakerReference sneakerReference) : base(string.Empty, Method.PATCH)
		{
			AddJsonBody(sneakerReference);
		}
	}
}
