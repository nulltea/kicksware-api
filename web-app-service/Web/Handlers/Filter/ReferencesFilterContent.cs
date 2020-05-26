using System;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Model;
using Core.Model.Parameters;
using Core.Reference;
using Web.Data.Catalog;

namespace Web.Handlers.Filter
{
	public class ReferencesFilterContent : FilterContentBuilder<SneakerReference>
	{
		public override void ConfigureFilter(IFilteredModel<SneakerReference> model)
		{
			if (AdditionalParams != null && AdditionalParams.TryGetValue("brandID", out var brandID))
			{
				model.AddHiddenFilterGroup("brandname", ExpressionType.Equal)
					.AssignParameter(Convert.ToString(brandID), Convert.ToString(brandID));
			}
			else if (AdditionalParams != null && AdditionalParams.TryGetValue("modelID", out var modelID))
			{
				model.AddHiddenFilterGroup("basemodel", ExpressionType.Equal)
					.AssignParameter(Convert.ToString(modelID), Convert.ToString(modelID));
			}
			else
			{
				model.AddFilterGroup("Brand", "brandname")
					.AssignParameters(Catalog.SneakerBrandsList);
			}
			model.AddForeignFilterGroup<SneakerProduct>("Size", "size")
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
			model.AddForeignFilterGroup<SneakerProduct>("Condition", "condition")
				.AssignParameters(typeof(SneakerCondition));
		}

		public override void ConfigureSorting(IFilteredModel<SneakerReference> model)
		{
			model.AddSortParameters(criterion => criterion switch
			{
				SortCriteria.Popular => new SortParameter(criterion, ""),
				SortCriteria.Newest => new SortParameter(criterion, "releasedate"),
				SortCriteria.Featured => new SortParameter(criterion, ""),
				SortCriteria.PriceFromLow => new SortParameter(criterion, "price", SortDirection.Ascending),
				SortCriteria.PriceFromHigh => new SortParameter(criterion, "price"),
				_ => throw new ArgumentException("No such sort criteria")
			});
			model.ChooseSortParameter("newest");
		}
	}
}