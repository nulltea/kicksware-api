using System;
using System.Collections.Generic;
using System.Linq;
using Core.Entities.Reference;
using Core.Model.Parameters;
using Core.Pattern;
using Core.Reference;

namespace Infrastructure.Pattern
{
	public class ReferencePredicateBuilder : IQueryPredicateBuilder<SneakerReference>
	{
		private List<FilterGroup> _queryGroups;

		public void SetQueryArguments(List<FilterGroup> groups) => _queryGroups = groups;

		public Func<SneakerReference, bool> Build()
		{
			return null;
		}
	}
}