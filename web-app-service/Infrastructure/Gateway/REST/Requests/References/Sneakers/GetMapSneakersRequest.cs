namespace Infrastructure.Gateway.REST.References.Sneakers
{
	public class GetMapSneakerReferencesRequest : BaseSneakerReferenceRequest
	{
		public GetMapSneakerReferencesRequest(object map) : base("/map")
		{
			AddJsonBody(map);
		}
	}
}