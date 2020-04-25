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
		public string Caption { get; }

		public string GroupID => (Caption ?? Property).ToLower();

		public FilterProperty Property { get; }

		public ExpressionType ExpressionType { get; }

		public List<FilterParameter> CheckedParameters => this.Where(param => param.Checked).ToList();

		public bool Hidden { get; set; } = false;

		public  string Description { get; }

		public FilterGroup(string caption, FilterProperty property, ExpressionType expressionType = ExpressionType.In, string description = default)
		{
			Caption = caption;
			Property = property;
			ExpressionType = expressionType;
			Description = description;
		}

		public FilterGroup(FilterProperty property, ExpressionType expressionType = ExpressionType.In)
		{
			Property = property;
			ExpressionType = expressionType;
		}

		public FilterGroup AssignParameter(FilterParameter parameter) => Add(parameter);

		public FilterGroup AssignParameter(string caption, object value,
											ExpressionType expressionType = ExpressionType.Equal,
											string description = default) =>
			Add(new FilterParameter(caption, value, expressionType, description));

		public FilterGroup AssignParameter(object value,
											ExpressionType expressionType = ExpressionType.Equal) =>
			Add(new FilterParameter(value, expressionType));

		public FilterGroup AssignParameters(params FilterParameter[] parameters) => AddRange(parameters);

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

		public FilterGroup AssignParameters<T>(IEnumerable<T> objects, Func<T, FilterParameter> selector) =>
			AddRange(objects.Select(selector));

		public FilterGroup AssignParameters(Dictionary<string, object> map) =>
			AddRange(map.Select(kvp => new FilterParameter(kvp.Key, kvp.Value)));


		public FilterGroup AssignParameters(IEnumerable<string> values)
		{
			AddRange(values.Select(value => new FilterParameter(value, value)));
			return this;
		}

		private new FilterGroup Add(FilterParameter parameter)
		{
			if (Hidden) parameter.Checked = true;
			base.Add(parameter);
			return this;
		}

		private new FilterGroup AddRange(IEnumerable<FilterParameter> collection)
		{
			var filterParameters = collection.ToList();
			if (Hidden) filterParameters.ForEach(param => param.Checked = true);
			base.AddRange(filterParameters);
			return this;
		}
	}
}