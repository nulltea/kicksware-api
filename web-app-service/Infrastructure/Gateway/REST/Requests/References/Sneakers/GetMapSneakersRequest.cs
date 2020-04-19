using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class GetMapSneakerReferencesRequest : BaseSneakerReferenceRequest
	{
		public GetMapSneakerReferencesRequest(Dictionary<string, object> map) : base("/map?", Method.POST)
		{
			AddJsonBody(map);
		}

		public GetMapSneakerReferencesRequest(object map) : base("/map", Method.POST)
		{
			AddJsonBody(map);
		}
	}
}