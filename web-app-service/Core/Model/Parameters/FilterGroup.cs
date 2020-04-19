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

		public string Property { get; set; }

		public ExpressionType ExpressionType { get; set; }

		public FilterGroup(string caption, string property, ExpressionType expressionType = ExpressionType.In, string description = default)
		{
			Caption = caption;
			Property = property;
			ExpressionType = expressionType;
			Description = description;
		}

		public void AssignParameter(FilterParameter parameter) => Add(parameter);

		public void AssignParameters(params FilterParameter[] parameters) => AddRange(parameters);

		public void AssignParameters(Type type)
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
		}

		public void AssignParameters<T>(IEnumerable<T> objects, Func<T, FilterParameter> selector) =>
			AddRange(objects.Select(selector));


		public void AssignParameters(Dictionary<string, object> map) =>
			AddRange(map.Select(kvp => new FilterParameter(kvp.Key, kvp.Value)));


		public void AssignParameters(IEnumerable<string> values) =>
			AddRange(values.Select(value => new FilterParameter(value, value)));
	}
}