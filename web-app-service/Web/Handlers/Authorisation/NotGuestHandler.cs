using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Extension;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Routing;
using Web.Handlers.Users;

namespace Web.Handlers.Authorisation
{
	public class NotGuestHandler : AuthorizationHandler<NotGuestRequirement>
	{
		private UserManager<User> _userManager;

		private IHttpContextAccessor _accessor;

		private LinkGenerator _linkGenerator;

		public NotGuestHandler(UserManager<User> userManager, IHttpContextAccessor accessor,
								LinkGenerator linkGenerator)
		{
			_userManager = userManager;
			_accessor = accessor;
			_linkGenerator = linkGenerator;
		}

		protected override async Task HandleRequirementAsync(AuthorizationHandlerContext context,
															NotGuestRequirement requirement)
		{
			if (context.User.IsInRole(UserRole.Guest.GetEnumMemberValue()))
			{
				AccessDenied(context, requirement);
				return;
			}

			var user = await _userManager.GetUserAsync(context.User);
			if (user is null || !user.Confirmed)
			{
				AccessDenied(context, requirement);
				return;
			}

			context.Succeed(requirement);
		}

		private void AccessDenied(AuthorizationHandlerContext context, IAuthorizationRequirement requirement)
		{
			var httpContext = _accessor.HttpContext;
			httpContext.Items.Add(
				AuthLockedResponse.AuthLockedKey,
				new AuthLockedResponse(httpContext.Request.Path)
			);
			context.Succeed(requirement);
		}
	}
}