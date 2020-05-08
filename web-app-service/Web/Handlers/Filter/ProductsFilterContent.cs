using System;
using Core.Entities.Products;
using Core.Model;
using Core.Model.Parameters;
using Core.Reference;
using Web.Data.Catalog;

namespace Web.Handlers.Filter
{
	public class ProductsFilterContent : FilterContentBuilder<SneakerProduct>
	{
		public override void ConfigureFilter(IFilteredModel<SneakerProduct> model)
		{
			if (AdditionalParams != null && AdditionalParams.TryGetValue("referenceId", out var referenceId)
										&& !string.IsNullOrEmpty(referenceId?.ToString()))
			{
				model.AddHiddenFilterGroup("ReferenceID",  ExpressionType.Equal)
					.AssignParameter(Convert.ToString(referenceId));
			}
			else
			{
				model.AddFilterGroup("Brand", "brandname")
					.AssignParameters(Catalog.SneakerBrandsList);
			}
			model.AddFilterGroup("Size", "size")
				.AssignParameters(Catalog.SneakerSizesList, size =>
					new FilterParameter(size.Europe.ToString("#.#"), size));
			model.AddFilterGroup("Color", "color", ExpressionType.Or)
				.AssignParameters(Catalog.FilterColors, color =>
					new FilterParameter(color.Name, color.Name.ToUpper(), ExpressionType.Regex) {SourceValue = color});
			model.AddFilterGroup("Price", "price", ExpressionType.And)
				.AssignParameters(
					new FilterParameter("Price (min)", 0, ExpressionType.GreaterThanOrEqual)
						{Checked = true, ImmutableValue = false},
					new FilterParameter("Price (max)", 1000, ExpressionType.LessThanOrEqual)
						{Checked = true, ImmutableValue = false}
				);
			model.AddFilterGroup("Condition", "condition")
				.AssignParameters(typeof(SneakerCondition));
		}

		public override void ConfigureSorting(IFilteredModel<SneakerProduct> model)
		{
			model.AddSortParameters(criterion => criterion switch
			{
				SortCriteria.Popular => new SortParameter(criterion, "likes"),
				SortCriteria.Newest => new SortParameter(criterion, "AddedAt"),
				SortCriteria.Featured => new SortParameter(criterion, "likes"),
				SortCriteria.PriceFromLow => new SortParameter(criterion, "price", SortDirection.Ascending),
				SortCriteria.PriceFromHigh => new SortParameter(criterion, "price"),
				_ => throw new ArgumentException("No such sort criteria")
			});
			model.ChooseSortParameter("newest");
		}
	}
}