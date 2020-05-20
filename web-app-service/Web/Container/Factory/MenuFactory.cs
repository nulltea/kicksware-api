using System;
using Infrastructure.Gateway.REST.Client;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Web.Config;
using Web.Handlers.Menu;

namespace Web.Container.Factory
{
	public partial class ServiceFactory
	{
		public static MenuBuilder<ShopMenuContent> ProvideShopMenuBuilder(IServiceProvider provider)
		{
			var options = provider.GetService<IOptions<AppSettings>>();
			if (options?.Value is null) throw new NullReferenceException("App settings not configured, although needed to build shop menu content");

			var builder = new MenuBuilder<ShopMenuContent>(options.Value.ShopMenuContentPath);

			return builder;
		}
	}
}