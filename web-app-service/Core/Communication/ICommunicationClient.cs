using System.Threading.Tasks;

namespace Core.Communication
{
	public interface ICommunicationClient
	{
		void Authenticate();

		void Request(ICommunicationRequest request);

		Task RequestAsync(ICommunicationRequest request);

		T Request<T>(ICommunicationRequest request) where T : class;

		Task<T> RequestAsync<T>(ICommunicationRequest request) where T : class;
	}
}