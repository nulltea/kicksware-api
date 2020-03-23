using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Reflection;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;
using web_app_service.Wizards;

namespace web_app_service.Utils.Helpers
{
	public static partial class CustomHelpers
	{
		public static IEnumerable<IHtmlContent> HiddenForWizardStep<TModel, TValue>(this IHtmlHelper<TModel> helper, WizardStep step, Expression<Func<TModel, TValue>> expression, IModelExpressionProvider m)
		{
			var model = m.CreateModelExpression(helper.ViewData, expression).Model;
			var modelDictionary = model.GetType()
				.GetProperties(BindingFlags.Instance | BindingFlags.Public)
				.Where(prop => !step.FormProperties.Contains(prop.Name))
				.ToDictionary(prop => prop.Name, prop => prop.GetValue(model, null));
			foreach (var (property, value) in modelDictionary)
			{
				yield return helper.Hidden(property, value);
			}
		}

		public static IEnumerable<IHtmlContent> HiddenForWizardStep(this IHtmlHelper helper, object model, WizardStep step)
		{
			var modelDictionary = model.GetType()
				.GetProperties(BindingFlags.Instance | BindingFlags.Public)
				.Where(prop => !step.FormProperties.Contains(prop.Name))
				.ToDictionary(prop => prop.Name, prop => prop.GetValue(model, null));
			foreach (var (property, value) in modelDictionary)
			{
				yield return helper.Hidden(property, value);
			}
		}
	}
}