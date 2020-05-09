using System;
using System.Collections.Generic;
using Core.Model.Parameters;

namespace Core.Pattern
{
	public interface IQueryPredicateBuilder<in T>
	{
		void SetQueryArguments(List<FilterGroup> groups);

		Func<T, bool> Build();
	}
}