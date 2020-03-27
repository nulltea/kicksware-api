using Microsoft.AspNetCore.Mvc;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		public ActionResult Details(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(0, model);

			return this.ViewStep(2, model);
		}
	}
}