using Core.Reference;

namespace Core.Gateway
{
	public class RequestParams
	{
		public int? Limit { get; set; }

		public int? Offset { get; set; }

		public string SortBy { get; set; }

		public SortDirection? SortDirection { get; set; }
	}
}