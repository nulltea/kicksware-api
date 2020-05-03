using System;
using System.Text.RegularExpressions;
using Core.Extension;
using Core.Reference;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace Core.Model.Parameters
{
	public class FilterParameter
	{
		public string RenderId => $"filter-control-{new Regex("[\\n\\t;,.\\s()\\/]").Replace((Caption ?? Convert.ToString(Value))!, "_").ToLower()}";

		public string Caption { get; }

		public string Description { get; }

		public object Value
		{
			get => _value;
			set => _value = ImmutableValue ? _value : value;
		}
		private object _value;

		public bool Checked { get; set; }

		public object SourceValue { get; set; }

		public ExpressionType ExpressionType { get; }

		public bool ImmutableValue { get; set; } = true;

		public FilterParameter(string caption, object value, ExpressionType expressionType = ExpressionType.Equal, string description=default)
		{
			Caption = caption;
			_value = value;
			ExpressionType = expressionType;
			Description = description;
		}

		public FilterParameter(object value, ExpressionType expressionType = ExpressionType.Equal)
		{
			_value = value;
			ExpressionType = expressionType;
			Checked = true;
		}

		public T GetSourceValue<T>() => (T)SourceValue;

		public IHtmlContent RenderCheckbox(object attributes = default)
		{
			var content = new TagBuilder("div");
			content.AddCssClass("filter-checkbox");

			var checkbox = new TagBuilder("input");
			checkbox.Attributes["type"] = "checkbox";
			checkbox.Attributes["id"] = RenderId;
			checkbox.Attributes["value"] = Convert.ToString(Value);
			if (attributes != default) checkbox.MergeAttributes(attributes.ToMap());
			if (Checked) checkbox.Attributes["checked"] = "true";
			content.InnerHtml.AppendHtml(checkbox);

			var label = new TagBuilder("label");
			label.Attributes["for"] = RenderId;
			label.InnerHtml.Append(Caption);
			content.InnerHtml.AppendHtml(label);

			return content;
		}

		public IHtmlContent RenderInput(object attributes = default)
		{
			var input = new TagBuilder("input");
			input.Attributes["type"] = "text";
			input.Attributes["id"] = RenderId;
			input.Attributes["value"] = Convert.ToString(Value);
			input.MergeAttributes(attributes.ToMap(), true);
			return input;
		}

		public IHtmlContent RenderHidden()
		{
			var hidden = new TagBuilder("input");
			hidden.Attributes["type"] = "hidden";
			hidden.Attributes["id"] = RenderId;
			hidden.Attributes["value"] = Convert.ToString(Value);
			return hidden;
		}
	}
}