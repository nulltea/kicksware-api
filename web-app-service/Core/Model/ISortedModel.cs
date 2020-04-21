using System;
using System.Collections.Generic;
using System.Linq;
using Core.Model.Parameters;
using Core.Reference;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Core.Model
{
	public interface ISortedModel
	{
		List<SortParameter> SortParameters { get; set; }

		SortParameter ChosenSorting { get; set; }

		SortParameter ChooseSortParameter(string value);

		void AddSortParameters(params SortParameter[] parameters);

		void AddSortParameters(Func<SortCriteria, SortParameter> selector);

		List<SelectListItem> FormSortSelect();
	}
}