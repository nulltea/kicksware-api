using Core.Entities.Products;
using Core.Entities.Users;
using Infrastructure.Gateway.REST.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class DeleteUserRequest : UserBaseRequest
	{
		public DeleteUserRequest(User user) : this(user.UniqueID) { }

		public DeleteUserRequest(string userID) : base("/{userID}", Method.DELETE)
		{
			AddParameter("userID", userID, ParameterType.UrlSegment);
		}
	}
}