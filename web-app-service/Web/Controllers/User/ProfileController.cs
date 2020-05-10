using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace Web.Controllers
{
	public class ProfileController : Controller
	{
		public async Task<IActionResult> Profile()
		{
			return View();
		}
	}
}