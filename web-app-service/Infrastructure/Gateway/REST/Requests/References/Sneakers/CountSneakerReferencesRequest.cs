using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class CountSneakerReferencesRequest : BaseSneakerReferenceRequest
	{
		public CountSneakerReferencesRequest() : base("/count") { }

		public CountSneakerReferencesRequest(Dictionary<string, object> query) : base("/count", Method.POST)
		{
			AddJsonBody(query);
		}

		public CountSneakerReferencesRequest(object query) : base("/count", Method.POST)
		{
			if (query != default) AddJsonBody(query);
		}
	}
}