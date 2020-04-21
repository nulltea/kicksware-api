using Core.Reference;

namespace Core.Gateway
{
	public class RequestParams
	{
		public int? TakeCount { get; set; }

		public int? SkipOffset { get; set; }

		public string SortBy { get; set; }

		public SortDirection? SortDirection { get; set; }
	}
}