using System.Collections;
using System.Collections.Generic;
using RestSharp;

namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class GetQueriedSneakerReferencesRequest : BaseSneakerReferenceRequest
	{
		public GetQueriedSneakerReferencesRequest(IEnumerable<string> idCodes) : base("/query")
		{
			foreach (var id in idCodes)
			{
				AddParameter("sneakerId", id, ParameterType.QueryString);
			}
		}
	}
}