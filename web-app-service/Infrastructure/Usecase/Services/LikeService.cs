using System.Threading.Tasks;
using Core.Gateway;
using Core.Services.Interactive;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Interact;

namespace Infrastructure.Usecase
{
	public class LikeService : ILikeService
	{
		private readonly IGatewayClient<IGatewayRestRequest> _client;

		public LikeService(IGatewayClient<IGatewayRestRequest> client) => _client = client;

		public void Like(string entityID) => _client.Request(new LikeRequest(entityID));

		public void Unlike(string entityID) => _client.Request(new UnlikeRequest(entityID));

		public Task LikeAsync(string entityID) => _client.RequestAsync(new LikeRequest(entityID));

		public Task UnlikeAsync(string entityID) => _client.RequestAsync(new UnlikeRequest(entityID));
	}
}