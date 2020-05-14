using System.Security.Claims;
using Core.Entities.Users;
using Core.Extension;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;

namespace Web.Handlers.Authentication
{
	public class MiddlewareSignInManager : SignInManager<User>
	{
		public MiddlewareSignInManager(UserManager<User> userManager,
										IHttpContextAccessor contextAccessor,
										IUserClaimsPrincipalFactory<User> claimsFactory,
										IOptions<IdentityOptions> optionsAccessor,
										ILogger<SignInManager<User>> logger,
										IAuthenticationSchemeProvider schemes,
										IUserConfirmation<User> confirmation)
			: base(userManager, contextAccessor, claimsFactory, optionsAccessor, logger, schemes, confirmation) { }

		public override bool IsSignedIn(ClaimsPrincipal principal)
		{
			if (principal is null || !principal.HasClaim(c => c.Type == ClaimTypes.NameIdentifier)) return false;

			if (principal.IsInRole(UserRole.Guest.GetEnumMemberValue())) return false;

			return true;
		}
	}
}