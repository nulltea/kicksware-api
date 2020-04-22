using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Mvc;

namespace Web.Controllers
{
	public class AuthController : Controller
	{
		public IActionResult Facebook()
		{
			var authProperties = new AuthenticationProperties
			{
				RedirectUri = Url.Action("Index", "Home")
			};
			return Challenge(authProperties, "Facebook");
		}

		public IActionResult Google()
		{
			var authProperties = new AuthenticationProperties
			{
				RedirectUri = Url.Action("Index", "Home")
			};
			return Challenge(authProperties, "Google");
		}
	}
}