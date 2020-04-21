using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using Core.Reference;

namespace Core.Model.Parameters
{
	public class FilterGroup : List<FilterParameter>
	{
		public string Caption { get; set; }

		public  string Description { get; set; }

		public FilterProperty Property { get; set; }

		public ExpressionType ExpressionType { get; set; }

		public List<FilterParameter> CheckedParameters => this.Where(param => param.Checked).ToList();

		public FilterGroup(string caption, FilterProperty property, ExpressionType expressionType = ExpressionType.In, string description = default)
		{
			Caption = caption;
			Property = property;
			ExpressionType = expressionType;
			Description = description;
		}

		public FilterGroup AssignParameter(FilterParameter parameter)
		{
			Add(parameter);
			return this;
		}

		public FilterGroup AssignParameters(params FilterParameter[] parameters)
		{
			AddRange(parameters);
			return this;
		}

		public FilterGroup AssignParameters(Type type)
		{
			var names = Enum.GetNames(type);
			foreach (var name in names)
			{
				var field = type.GetField(name);
				var displayAttr = field.GetCustomAttribute<DisplayAttribute>(true);
				var memberAttr = field.GetCustomAttribute<EnumMemberAttribute>(true);
				if (displayAttr is null || memberAttr is null) continue;
				Add(new FilterParameter(displayAttr.Name, memberAttr.Value, description: displayAttr.GetDescription()));
			}

			return this;
		}

		public FilterGroup AssignParameters<T>(IEnumerable<T> objects, Func<T, FilterParameter> selector)
		{
			AddRange(objects.Select(selector));
			return this;
		}

		public FilterGroup AssignParameters(Dictionary<string, object> map)
		{
			AddRange(map.Select(kvp => new FilterParameter(kvp.Key, kvp.Value)));
			return this;
		}



		public FilterGroup AssignParameters(IEnumerable<string> values)
		{
			AddRange(values.Select(value => new FilterParameter(value, value)));
			return this;
		}

	}
}