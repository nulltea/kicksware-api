using Core.Entities;
using Microsoft.AspNetCore.Mvc;

namespace Web.Wizards
{
	public interface IEntityWizard<in T> where T : class, IBaseEntity
	{
		ActionResult InitWizard(Controller controller, T model);

		ActionResult NextStep(Controller controller, T model);

		ActionResult PreviousStep(Controller controller, T model);
	}
}