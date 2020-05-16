using System;
using System.Linq;
using System.Net;
using System.Runtime.Serialization;
using System.Security.Authentication;
using System.Threading.Tasks;
using Core.Constants;
using Core.Entities.Users;
using Core.Extension;
using Core.Gateway;
using Infrastructure.Serializers;
using RestSharp;
using RestSharp.Authenticators;

namespace Infrastructure.Gateway.REST.Client
{
	public class RestfulClient : RestClient, IGatewayClient<IGatewayRestRequest>
	{
		public RestfulClient() : base(Constants.GatewayBaseUrl)
		{
			UseSerializer(() => new JsonRestSerializer());
		}

		public void Authenticate(AuthToken token)
		{
			Authenticator = new JwtAuthenticator(token);
		}

		public bool Request(IGatewayRestRequest request)
		{
			var response = Execute(ApplyRequestParams(request));

			return HandleRequestStatus(response) && response.IsSuccessful;
		}

		public async Task<bool> RequestAsync(IGatewayRestRequest request)
		{
			var response = await ExecuteAsync(ApplyRequestParams(request));
			try
			{
				HandleRequestStatus(response);
			}
			catch
			{
				return false;
			}
			return response.IsSuccessful;
		}

		public T Request<T>(IGatewayRestRequest request)
		{
			var response = Execute<T>(ApplyRequestParams(request));
			if (!HandleRequestStatus<T>(response, out var data)) return default;
			return data;
		}

		public async Task<T> RequestAsync<T>(IGatewayRestRequest request)
		{
			var response = await ExecuteAsync<T>(ApplyRequestParams(request));
			if (!HandleRequestStatus<T>(response, out var data)) return default;
			return data;
		}

		private bool HandleRequestStatus<T>(IRestResponse<T> response, out T data)
		{
			switch (response.StatusCode)
			{
				case HttpStatusCode.OK:
					data = response.Data ?? Activator.CreateInstance<T>();
					return true;
				case 0:
				case HttpStatusCode.NotFound:
				case HttpStatusCode.NoContent:
				case HttpStatusCode.NotModified:
				case HttpStatusCode.NotImplemented:
					data = Activator.CreateInstance<T>();
					return true;
				case HttpStatusCode.Unauthorized:
					throw new AuthenticationException(response.Content);
				default:
					throw new Exception(response.Content);
			}
		}

		private bool HandleRequestStatus(IRestResponse response)
		{
			switch (response.StatusCode)
			{
				case HttpStatusCode.OK:
					return true;
				case 0:
				case HttpStatusCode.NotFound:
				case HttpStatusCode.NoContent:
				case HttpStatusCode.NotModified:
				case HttpStatusCode.NotImplemented:
					return false;
				case HttpStatusCode.Unauthorized:
					throw new AuthenticationException(response.Content);
				default:
					throw new Exception(response.Content);
			}
		}

		private static IGatewayRestRequest ApplyRequestParams(IGatewayRestRequest request)
		{
			var parameters = request.RequestParams?.ToMap().Where(kvp => kvp.Value != null);
			if (parameters is null) return request;
			foreach (var pair in parameters)
			{
				var (key, value) = pair;
				if (value is Enum enumVal) value = enumVal.GetEnumAttribute<EnumMemberAttribute>()?.Value ?? value;
				request.AddParameter(key.ToLower(), value, ParameterType.QueryString);
			}
			return request;
		}
	}
}