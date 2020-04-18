using System;

namespace Core.Exceptions
{
	public class NextPageNotValidException : Exception
	{
		public override string Message => "Next page does not exists";
	}

	public class PreviousPageNotValidException : Exception
	{
		public override string Message => "Previous page does not exists";
	}

	public class PageNotValidException : Exception
	{
		private int Page { get; }

		public PageNotValidException(int page) => Page = page;

		public override string Message => $"Page by number {Page} does not exists";
	}
}