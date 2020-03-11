using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST
{
	public interface IGatewayRestRequest : IRestRequest, IGatewayRequest { }
}