using System;
using System.Collections.Generic;
using System.Linq;
using Core.Attributes;
using Core.Extension;
using Core.Model.Parameters;
using Core.Reference;

namespace Infrastructure.Pattern
{
	public class QueryRecourseBuilder
	{
		private List<FilterGroup> _queryGroups;

		public void SetQueryArguments(FilterGroup group) => _queryGroups = new List<FilterGroup> {group};

		public void SetQueryArguments(List<FilterGroup> groups) => _queryGroups = groups;

		public void SetQueryArguments(List<FilterParameter> parameters, FilterProperty property, ExpressionType expressionType = ExpressionType.Or) => _queryGroups = new List<FilterGroup>
		{
			new FilterGroup(property, property, expressionType).AssignParameters(parameters.ToArray())
		};

		public Dictionary<string, object> Build()
		{
			var resultQuery = new Dictionary<string, object>();
			foreach (var filterGroup in _queryGroups)
			{
				if (filterGroup.Property.IsForeignEntity) //handle subservice entity query
				{
					var subgroup = new FilterGroup(filterGroup.Caption, filterGroup.Property.FieldName, filterGroup.ExpressionType)
							.AssignParameters(filterGroup.CheckedParameters.ToArray());
					if (!BuildForGroup(subgroup, out var subqueryPair)) continue;
					var subquery = new Dictionary<string, object>
					{
						{subqueryPair.property, subqueryPair.query}
					};
					var subservice = $"*/{filterGroup.Property.ForeignResource}";
					if (!resultQuery.TryAdd(subservice, subquery))
					{
						if (resultQuery[subservice] is Dictionary<string, object> existedQuery)
						{
							existedQuery.TryAdd(subqueryPair.property, subquery[subqueryPair.property]);
							resultQuery[subservice] = existedQuery;
						}
					}
					continue;
				}

				if (!BuildForGroup(filterGroup, out var queryPair)) continue;

				resultQuery.TryAdd(queryPair.property, queryPair.query);
			}
			return resultQuery;
		}

		private bool BuildForGroup(FilterGroup filterGroup, out (string property, object query) resultQuery)
		{
			resultQuery = default;
			var checkedParams = filterGroup.CheckedParameters;
			if (!checkedParams.Any()) return false;

			var multiply = checkedParams.Count > 1;
			if (multiply)
			{
				var groupExpression = filterGroup.ExpressionType.GetEnumAttribute<QueryExpressionAttribute>();
				switch (groupExpression.Target)
				{
					case ExpressionTarget.Group:
					case ExpressionTarget.Both:
					{
						var listQuery = new Dictionary<string, object>
						{
							{ groupExpression.OperatorSyntax, checkedParams.Select(param => FormatQueryValue(groupExpression, param.Value)) }
						};
						resultQuery = (filterGroup.Property, listQuery);
						break;
					}
					case ExpressionTarget.Each:
					{
						var eachParamQuery = new List<Dictionary<string, object>>();
						foreach (var param in checkedParams)
						{
							if (param.ExpressionType == ExpressionType.Equal)
							{
								eachParamQuery.Add(new Dictionary<string, object>{{filterGroup.Property, param.Value}});
							}
							else
							{
								var nodeExpression = param.ExpressionType.GetEnumAttribute<QueryExpressionAttribute>();
								var operatorCondition = new Dictionary<string, object>{{nodeExpression.OperatorSyntax, FormatQueryValue(nodeExpression, param.Value)}};
								eachParamQuery.Add(new Dictionary<string, object>{{filterGroup.Property, operatorCondition}});
							}
						}
						resultQuery = (groupExpression.OperatorSyntax, eachParamQuery);
						break;
					}
					case ExpressionTarget.Node:
					{
						var eachParamQuery = new Dictionary<string, object>();
						foreach (var param in checkedParams)
						{
							var nodeExpression = param.ExpressionType.GetEnumAttribute<QueryExpressionAttribute>();
							eachParamQuery.Add(nodeExpression.OperatorSyntax, FormatQueryValue(nodeExpression, param.Value));
						}
						resultQuery = (filterGroup.Property, eachParamQuery);
						break;
					}
					default:
						throw new ArgumentOutOfRangeException(nameof(groupExpression.Target));
				}

				return true;
			}

			var singleNode = checkedParams.First();
			if (singleNode.ExpressionType != ExpressionType.Equal)
			{
				var nodeOperator = singleNode.ExpressionType.GetEnumAttribute<QueryExpressionAttribute>();
				var operatorCondition = new Dictionary<string,object>{{nodeOperator.OperatorSyntax, FormatQueryValue(nodeOperator, singleNode.Value)}};
				resultQuery = (filterGroup.Property, operatorCondition);
				return true;
			}
			resultQuery = (filterGroup.Property, singleNode.Value);
			return true;
		}

		private static object FormatQueryValue(QueryExpressionAttribute expAttr, object value)
		{
			if (string.IsNullOrEmpty(expAttr.ValueWrapperFormat)) return value;

			return string.Format(expAttr.ValueWrapperFormat, value);
		}
	}
}