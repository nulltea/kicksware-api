using System.Collections;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class GetQueriedSneakersRequest : BaseSneakersListRequest
	{
		public GetQueriedSneakersRequest(object queryObject) : base(string.Empty)
		{
			if (queryObject is IEnumerable)
			{
				Resource += "/{sneakerId}/";
				AddParameter("sneakerId", queryObject, ParameterType.QueryString);
				return;
			}
			AddObject(queryObject);
		}
	}
}