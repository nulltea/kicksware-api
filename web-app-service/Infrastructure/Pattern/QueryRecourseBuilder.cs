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

		public void SetQueryArguments(List<FilterGroup> groups) => _queryGroups = groups;

		public Dictionary<string, object> Build()
		{
			var resultQuery = new Dictionary<string, object>();
			foreach (var filterGroup in _queryGroups)
			{
				var checkedParams = filterGroup.Where(param => param.Checked).ToList();
				if (!checkedParams.Any()) continue;

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
								{ groupExpression.OperatorSyntax, checkedParams.Select(param =>  FormatQueryValue(groupExpression, param.Value)) }
							};
							resultQuery.TryAdd(filterGroup.Property, listQuery);
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
							resultQuery.TryAdd(groupExpression.OperatorSyntax, eachParamQuery);
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
							resultQuery.TryAdd(filterGroup.Property, eachParamQuery);
							break;
						}
						default:
							throw new ArgumentOutOfRangeException(nameof(groupExpression.Target));
					}
					continue;
				}

				var singleNode = checkedParams.First();
				if (singleNode.ExpressionType != ExpressionType.Equal)
				{
					var nodeOperator = singleNode.ExpressionType.GetEnumAttribute<QueryExpressionAttribute>();
					resultQuery.Add(filterGroup.Property, new KeyValuePair<string,object>(nodeOperator.OperatorSyntax, FormatQueryValue(nodeOperator, singleNode.Value)));
					continue;
				}
				resultQuery.Add(filterGroup.Property, singleNode.Value);
			}
			return resultQuery;
		}

		private static object FormatQueryValue(QueryExpressionAttribute expAttr, object value)
		{
			if (string.IsNullOrEmpty(expAttr.ValueWrapperFormat)) return value;

			return string.Format(expAttr.ValueWrapperFormat, value);
		}
	}
}