using System;
using System.Collections.Generic;
using System.Linq;
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
					var groupOperator = Convert.ToString(filterGroup.ExpressionType.GetEnumMemberValue());
					//TODO OR & AND handle
					var listQuery = new Dictionary<string, object>
					{
						{groupOperator, checkedParams.Select(param => param.Value) }
					};
					resultQuery.TryAdd(filterGroup.Property, listQuery);
					continue;
				}

				var singleNode = checkedParams.First();
				if (singleNode.ExpressionType != ExpressionType.Equal)
				{
					var nodeOperator = Convert.ToString(singleNode.ExpressionType.GetEnumMemberValue());
					var otherQuery = new Dictionary<string, object>
					{
						{nodeOperator, singleNode.Value }
					};
					resultQuery.Add(filterGroup.Property, otherQuery);
					continue;
				}
				resultQuery.Add(filterGroup.Property, singleNode.Value);
			}
			resultQuery.Add("brandname", "Nike");
			return resultQuery;
		}
	}
}