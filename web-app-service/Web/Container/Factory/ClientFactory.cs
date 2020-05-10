using System;
using Core.Gateway;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Client;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Web.Utils.Extensions;

namespace Web.Container.Factory
{
	public static partial class ServiceFactory
	{
		public static RestfulClient ProvideRestClient(IServiceProvider serviceProvider)
		{
			if (serviceProvider is null) throw new ArgumentNullException(nameof(serviceProvider));

			var contextAccessor = serviceProvider.GetService<IHttpContextAccessor>();
			if (contextAccessor is null) throw new NullReferenceException(nameof(contextAccessor));

			var context = contextAccessor.HttpContext;
			if (context is null) throw new NullReferenceException(nameof(context));

			var client = new RestfulClient();

			var token = context.RetrieveCookieAuthToken();
			if (!string.IsNullOrEmpty(token)) client.Authenticate(token);

			return client;
		}
	}
}