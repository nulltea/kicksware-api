using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using Core.Constants;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;
using Microsoft.AspNetCore.Razor.TagHelpers;

namespace Web.Wizards
{
	public class WizardStep
	{
		public string View { get; set; }

		public string Icon { get; }

		public string Description { get; set; }

		public bool Active { get; set; }

		public bool Passed { get; set; }

		public string Name { get; set; }

		public string[] FormProperties;

		public WizardStep NextStep { get; set; }

		public WizardStep PreviousStep { get; set; }

		public WizardStep Next()
		{
			Passed = true;
			Active = false;
			if (NextStep is null) throw new Exception("Next step is not defined");
			NextStep.Active = true;
			return NextStep;
		}

		public WizardStep Previous()
		{
			Passed = Active = false;
			if (PreviousStep is null) throw new Exception("Previous step is not defined");
			PreviousStep.Active = true;
			return PreviousStep;
		}

		public WizardStep(string name, string view, string icon) => (Name, View, Icon) = (name, view, icon);

		public IHtmlContent RenderStep()
		{
			var span = new TagBuilder("span");
			span.InnerHtml.AppendHtml(File.ReadAllText(Path.Combine(Constants.ImagesPath, Icon)));

			var stepListItem = new TagBuilder("li");
			var stepListClass = Passed ? "passed" : Active ? "active" : string.Empty;
			if (!string.IsNullOrEmpty(stepListClass)) stepListItem.AddCssClass(stepListClass);
			stepListItem.InnerHtml.AppendHtml(span);

			return stepListItem;
		}
	}
}