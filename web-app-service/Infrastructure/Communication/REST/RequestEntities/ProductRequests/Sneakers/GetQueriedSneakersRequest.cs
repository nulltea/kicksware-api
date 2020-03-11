using System.Collections;
using System.Linq;
using RestSharp;

namespace Infrastructure.Communication.REST.ProductRequests.Sneakers
{
	public class GetQueriedSneakersRequest : BaseSneakersListRequest
	{
		public GetQueriedSneakersRequest(object queryObject) : base("")
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