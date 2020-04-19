namespace Core.Gateway
{
	public class RequestParams
	{
		public int? TakeCount { get; set; }

		public int? SkipOffset { get; set; }

		public bool? Pretty { get; set; }
	}
}