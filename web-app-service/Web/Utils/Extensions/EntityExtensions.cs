using System.Collections.Generic;
using System.Security.Claims;
using Core.Entities.Users;

namespace Web.Utils.Extensions
{
	public static class EntityExtensions
	{
		public static List<Claim> ExtractCredentials(this User user)
		{
			return new List<Claim>
			{
				new Claim(ClaimTypes.Email, user.Email),
				new Claim(ClaimTypes.Hash, user.PasswordHash)
			};
		}
	}
}