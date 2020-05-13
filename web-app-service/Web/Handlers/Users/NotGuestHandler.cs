using System.Linq;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Extension;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.Routing;
using Microsoft.AspNetCore.Routing;

namespace Web.Handlers.Users
{
	public class NotGuestHandler : AuthorizationHandler<NotGuestRequirement>
	{
		private UserManager<User> _userManager;

		private IHttpContextAccessor _accessor;

		private LinkGenerator _linkGenerator;

		public NotGuestHandler(UserManager<User> userManager, IHttpContextAccessor accessor, LinkGenerator linkGenerator)
		{
			_userManager = userManager;
			_accessor = accessor;
			_linkGenerator = linkGenerator;
		}

		protected override async Task HandleRequirementAsync(AuthorizationHandlerContext context,
															NotGuestRequirement requirement)
		{
			var httpContext = _accessor.HttpContext;
			if (context.User.IsInRole(UserRole.Guest.GetEnumMemberValue()))
			{
				context.Succeed(requirement);
				AccessDenied(httpContext);
				return;
			}

			var user = await _userManager.GetUserAsync(context.User);

			if (user is null || !user.Confirmed)
			{
				context.Succeed(requirement);
				AccessDenied(httpContext);
			}
			context.Succeed(requirement);
		}

		private void AccessDenied(HttpContext context)
		{
			context.Items.Add("locked", true);
			//context.Response.Redirect(_linkGenerator.GetPathByAction("AccessDenied", "Auth", new {fromAction = context.Request.Path}));
		}
	}
}