using Core.Entities.Products;
using Core.Entities.Users;
using Infrastructure.Gateway.REST.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.ProductRequests.Sneakers
{
	public class DeleteUserRequest : UserBaseRequest
	{
		public DeleteUserRequest(User user) : this(user.UserName) { }

		public DeleteUserRequest(string username) : base("/{username}", Method.DELETE)
		{
			AddParameter("username", username, ParameterType.UrlSegment);
		}
	}
}