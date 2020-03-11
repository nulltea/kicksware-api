using System.Threading.Tasks;
using Core.Communication;
using RestSharp;

namespace Infrastructure.Communication.REST.Client
{
	public class RestfulClient : RestClient, ICommunicationClient
	{
		public void Authenticate() {}

		public void Request(ICommunicationRequest request)
		{

		}

		public Task RequestAsync(ICommunicationRequest request)
		{
			return null;
		}

		public T Request<T>(ICommunicationRequest request) where T : class
		{
			return null;
		}

		public Task<T> RequestAsync<T>(ICommunicationRequest request) where T : class
		{
			return null;
		}
	}
}