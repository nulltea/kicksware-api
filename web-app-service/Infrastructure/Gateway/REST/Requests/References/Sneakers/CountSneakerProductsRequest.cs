namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class CountSneakerReferencesRequest : BaseSneakerReferenceRequest
	{
		public CountSneakerReferencesRequest(object queryObject) : base("/count")
		{
			AddObject(queryObject);
		}
	}
}