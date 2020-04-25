using System;
using System.Collections.Generic;
using Core.Entities;
using Core.Model.Parameters;
using Core.Reference;

namespace Core.Model
{
	public interface IFilteredModel<out TEntity> : IPagedModel<TEntity>, ISortedModel where TEntity : IBaseEntity
	{
		List<FilterGroup> FilterGroups { get; }

		List<FilterParameter> FilterParameters { get; }

		FilterGroup AddFilterGroup(string caption, FilterProperty property, ExpressionType expressionType = ExpressionType.In,
							string description = default);

		FilterGroup AddForeignFilterGroup<TForeignEntity>(string caption, string fieldName,
													ExpressionType expressionType = ExpressionType.In,
													string description = default);

		FilterGroup AddForeignFilterGroup(string caption, string fieldName, Type foreignEntity,
										ExpressionType expressionType = ExpressionType.In,
										string description = default);

		FilterGroup AddHiddenFilterGroup(FilterProperty property, ExpressionType expressionType = ExpressionType.In);

		FilterGroup GetFilterGroup(string name);

		void ApplyUserInputs(List<FilterInput> filterInputs);

		public void FetchFiltered();

		FilterGroup this[string groupName] { get; set; }
	}
}