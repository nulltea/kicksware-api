using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Reflection;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;
using Newtonsoft.Json;
using Web.Wizards;

namespace Web.Utils.Helpers
{
	public static partial class CustomHelpers
	{
		public static IEnumerable<IHtmlContent> HiddenForWizardStep<TModel, TValue>(this IHtmlHelper<TModel> helper, WizardStep step, Expression<Func<TModel, TValue>> expression, IModelExpressionProvider m)
		{
			var model = m.CreateModelExpression(helper.ViewData, expression).Model;

			foreach (var html in helper.HiddenForWizardStep(model, step)) yield return html;
		}

		public static IEnumerable<IHtmlContent> HiddenForWizardStep(this IHtmlHelper helper, object model, WizardStep step)
		{
			var modelDictionary = model.GetType()
				.GetProperties(BindingFlags.Instance | BindingFlags.Public)
				.Where(prop => !step.FormProperties.Contains(prop.Name))
				.ToDictionary(prop => prop.Name, prop => prop.GetValue(model, null));
			foreach (var (property, value) in modelDictionary)
			{
				if (value is IList list && value.GetType().IsGenericType)
				{
					for (var index = 0; index < list.Count; index++)
					{
						yield return helper.Hidden($"{property}[{index}]", list[index]);
					}

					continue;
				}
				yield return helper.Hidden(property, value);
			}
		}

		public static IEnumerable<IHtmlContent> HiddenForPartialModel(this IHtmlHelper helper, object model, params string[] exceptFields)
		{
			var modelDictionary = model.GetType()
				.GetProperties(BindingFlags.Instance | BindingFlags.Public)
				.Where(prop => !exceptFields.ToList()
					.Exists(f => string.Equals(f, prop.Name, StringComparison.CurrentCultureIgnoreCase)))
				.ToDictionary(prop => prop.Name, prop => prop.GetValue(model, null));
			foreach (var (property, value) in modelDictionary)
			{
				if (value is IList list && value.GetType().IsGenericType)
				{
					foreach (var valueItem in list)
					{
						yield return helper.Hidden(property, valueItem);
					}
					continue;
				}
				yield return helper.Hidden(property, value);
			}
		}
	}
}