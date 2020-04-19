using System.Collections.Generic;
using Core.Model.Parameters;

namespace Core.Pattern
{
	public interface IQueryRecourseBuilder
	{
		void SetQueryArguments(List<FilterGroup> groups);

		Dictionary<string, object> Build();
	}
}