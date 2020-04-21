using System;
using System.Collections.Generic;
using Core.Entities;
using Core.Model.Parameters;
using Core.Reference;

namespace Core.Model
{
	public interface IFilteredModel<T> : IPagedModel<T>, ISortedModel where T : IBaseEntity
	{
		List<FilterGroup> FilterGroups { get; }

		List<FilterParameter> FilterParameters { get; }

		FilterGroup AddFilterGroup(string caption, FilterProperty property, ExpressionType expressionType = ExpressionType.In,
							string description = default);

		FilterGroup AddForeignFilterGroup<TEntity>(string caption, string fieldName,
													ExpressionType expressionType = ExpressionType.In,
													string description = default);

		FilterGroup AddForeignFilterGroup(string caption, string fieldName, Type foreignEntity,
										ExpressionType expressionType = ExpressionType.In,
										string description = default);

		FilterGroup GetFilterGroup(string name);

		void ApplyUserInputs(Dictionary<string, (bool Checked, object Value)> filterInputs);

		public void FetchFiltered();

		FilterGroup this[string groupName] { get; set; }
	}
}