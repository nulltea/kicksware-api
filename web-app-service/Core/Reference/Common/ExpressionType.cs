using Core.Attributes;

namespace Core.Reference
{
	public enum ExpressionType
	{
		[QueryExpression(OperatorSyntax = "$and", Target = ExpressionTarget.Each)]
		And,

		[QueryExpression(OperatorSyntax = "$or", Target = ExpressionTarget.Each)]
		Or,

		[QueryExpression(OperatorSyntax = "$in", Target = ExpressionTarget.Group)]
		In,

		[QueryExpression(OperatorSyntax = "$nin", Target = ExpressionTarget.Group)]
		NotIn,

		[QueryExpression(OperatorSyntax = "$eq", Target = ExpressionTarget.Node)]
		Equal,

		[QueryExpression(OperatorSyntax = "$ne", Target = ExpressionTarget.Node)]
		NotEqual,

		[QueryExpression(OperatorSyntax = "$regex", Target = ExpressionTarget.Node, ValueWrapperFormat = ".*{0}.*")]
		Regex,

		[QueryExpression(OperatorSyntax = "$ge", Target = ExpressionTarget.Node)]
		GreaterThen,

		[QueryExpression(OperatorSyntax = "$gte", Target = ExpressionTarget.Node)]
		GreaterThanOrEqual,

		[QueryExpression(OperatorSyntax = "$le", Target = ExpressionTarget.Node)]
		LessThan,

		[QueryExpression(OperatorSyntax = "$lte", Target = ExpressionTarget.Node)]
		LessThanOrEqual,
	}
}