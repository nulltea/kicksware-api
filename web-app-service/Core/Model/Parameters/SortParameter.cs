using System;
using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;
using Core.Extension;
using Core.Reference;

namespace Core.Model.Parameters
{
	public class SortParameter
	{
		public string Caption { get; }

		public SortCriteria? Criteria { get; set; }

		public FilterProperty Property { get; set; }

		public SortDirection Direction { get; set; }

		public string RenderValue => (Criteria?.GetEnumAttribute<EnumMemberAttribute>()?.Value ?? Caption).ToLower();

		public SortParameter(SortCriteria criteria, FilterProperty property, SortDirection direction = SortDirection.Descending)
		{
			Criteria = criteria;
			Caption = criteria.GetEnumAttribute<DisplayAttribute>()?.Name;
			Property = property;
			Direction = direction;
		}

		public SortParameter(string caption, FilterProperty property, SortDirection direction = SortDirection.Descending)
		{
			Caption = caption;
			Property = property;
			Direction = direction;
		}
	}
}