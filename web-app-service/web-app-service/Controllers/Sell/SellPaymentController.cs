using System.Linq;
using Microsoft.AspNetCore.Mvc;
using web_app_service.Data.Reference_Data;
using web_app_service.Models;
using web_app_service.Wizards;

namespace web_app_service.Controllers
{
	public partial class SellController
	{
		public ActionResult Payment(SneakerProductViewModel model, bool rollback)
		{
			if (rollback) return this.ViewStep(2, model);
			if (model.ShippingInfo is null || !model.ShippingInfo.Any())
			{
				model.ShippingInfo = Catalog.DefaultShippingInfo;
			}
			return this.ViewStep(4, model);
		}
	}
}