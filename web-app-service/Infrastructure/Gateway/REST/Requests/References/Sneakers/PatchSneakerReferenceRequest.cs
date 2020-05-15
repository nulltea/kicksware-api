using System.IO;
using Core.Entities.Products;
using Core.Entities.References;
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
