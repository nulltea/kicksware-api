using System.Threading.Tasks;
using Core.Entities.Users;

namespace Core.Gateway
{
	public interface IGatewayClient<in T> where T : IGatewayRequest
	{
		void Authenticate(AuthToken token);

		bool Request(T request);

		Task<bool> RequestAsync(T request);

		TR Request<TR>(T request);

		Task<TR> RequestAsync<TR>(T request);
	}


}