using System;
using System.Reflection;
using Core.Attributes;
using Core.Gateway;
using Core.Reference;

namespace Core.Model.Parameters
{
	public class FilterProperty
	{
		public string FieldName { get; set; }

		public bool IsForeignEntity { get; set; } = false;

		public Type ForeignEntityType { get; set; }

		public string ForeignResource => ForeignEntityType.GetCustomAttribute<EntityServiceAttribute>()?.Resource;

		//public ExpressionType SubqueryExpression { get; set; } = ExpressionType.Exists;

		public static implicit operator FilterProperty(string field) => new FilterProperty(field);

		public static implicit operator string(FilterProperty property) => property.FieldName;

		public FilterProperty(string fieldName) => FieldName = fieldName;

		public FilterProperty(string fieldName, Type foreignEntity, bool isForeignEntity = true)
		{
			IsForeignEntity = isForeignEntity;
			FieldName = fieldName;
			ForeignEntityType = foreignEntity;
		}

		public override string ToString() => FieldName;
	}
}