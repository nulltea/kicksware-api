using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.Search
{
	public class SearchRequest : RestRequest, IGatewayRestRequest
	{
		public RequestParams RequestParams { get; set; }

		protected SearchRequest(string entity, string resource)
			: base($"search/{{entity}}{resource}", Method.GET)
		{
			AddParameter("entity", entity, ParameterType.UrlSegment);
		}
	}
}