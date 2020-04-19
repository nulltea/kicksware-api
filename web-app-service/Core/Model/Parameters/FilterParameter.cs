using Core.Reference;

namespace Core.Model.Parameters
{
	public class FilterParameter
	{
		public FilterParameter(string caption, object value, ExpressionType expressionType = ExpressionType.Equal, string description=default)
		{
			Caption = caption;
			Value = value;
			ExpressionType = expressionType;
			Description = description;
		}

		public string Caption { get; set; }

		public string Description { get; set; }

		public object Value { get; set; }

		public bool Checked { get; set; }

		public object SourceValue { get; set; }

		public ExpressionType ExpressionType { get; set; }
	}
}