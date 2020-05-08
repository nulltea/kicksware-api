using Microsoft.Extensions.Options;

namespace Web.Auth
{
	public class MiddlewareAuthPostConfigureOptions : IPostConfigureOptions<MiddlewareAuthOptions>
	{
		public void PostConfigure(string name, MiddlewareAuthOptions options)
		{

		}
	}
}