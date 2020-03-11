using System;
using System.Net;
using System.Security.Authentication;
using System.Threading.Tasks;
using Core.Constants;
using Core.Gateway;
using RestSharp;

namespace Infrastructure.Gateway.REST.Client
{
	public class RestfulClient : RestClient, IGatewayClient<IGatewayRestRequest>
	{
		internal RestfulClient() : base(Constants.GatewayBaseUrl) { }
		
		public void Authenticate() { }

		public bool Request(IGatewayRestRequest request)
		{
			var response = Execute(request);
			GuardUnsuccessfulRequest(request, response);
			return response.IsSuccessful;
		}

		public async Task<bool> RequestAsync(IGatewayRestRequest request)
		{
			var response = await ExecuteAsync(request);
			GuardUnsuccessfulRequest(request, response);
			return response.IsSuccessful;
		}

		public T Request<T>(IGatewayRestRequest request)
		{
			var response = Execute<T>(request);
			GuardUnsuccessfulRequest(request, response);

			if (response.StatusCode == HttpStatusCode.NotFound) return default;

			return response.Data;
		}

		public async Task<T> RequestAsync<T>(IGatewayRestRequest request)
		{
			var response = await ExecuteAsync<T>(request);
			GuardUnsuccessfulRequest(request, response);
			return response.Data;
		}

		private void GuardUnsuccessfulRequest(IGatewayRestRequest request, IRestResponse response)
		{
			if (request.Method == Method.GET && response.StatusCode == HttpStatusCode.NotFound) return;
			if (response.StatusCode == HttpStatusCode.Unauthorized) throw new AuthenticationException(response.Content);
			if (response.StatusCode != HttpStatusCode.OK) throw new Exception(response.Content);
		}
	}
}