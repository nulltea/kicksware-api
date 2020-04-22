using System;
using Microsoft.AspNetCore.Mvc;
using Web.Models;

namespace Web.Wizards
{
	public static class SneakerProductWizard
	{
		private static WizardSteps Steps(int active = 0) => new WizardSteps(new[]
		{
			new WizardStep("Search", "New/Search", "cloud_search.svg")
			{
				FormProperties = new string[]{}
			},
			new WizardStep("Details", "New/Details", "details.svg")
			{
				FormProperties = new[] { "BrandName", "ModelName", "Type", "Size", "Color", "Condition", "Description" }
			},
			new WizardStep("Photos", "New/Photos", "photo.svg")
			{
				FormProperties = new[] { "FormFiles", "Photos" }
			},

			new WizardStep("Payment", "New/Payment", "payment.svg")
			{
				FormProperties =  new[] {"Price", "Currency"}
			},
			new WizardStep("Shipping", "New/Shipping", "shipping.svg")
			{
				FormProperties =  new[] {"ShippingInfo"}
			},
			new WizardStep("Preview", "New/Preview", "preview.svg")
			{
				FormProperties =  new[] {"*"}
			},
		}, active);

		public static ActionResult ViewStep(this Controller controller, int stepIndex, SneakerProductViewModel model)
		{
			var steps = Steps(stepIndex);
			controller.ViewBag.Steps = steps;
			return controller.View(steps.ActiveStep.View, model);
		}
	}
}