using System;

namespace Core.Attributes
{
	[AttributeUsage(AttributeTargets.Field)]
	public class QueryExpressionAttribute : Attribute
	{
		public string OperatorSyntax { get; set; }

		public ExpressionTarget Target { get; set; } = ExpressionTarget.Both;

		public string ValueWrapperFormat { get; set; }
	}

	public enum ExpressionTarget
	{
		Group,

		Node,

		Both,

		Each,

		Subquery
	}
}