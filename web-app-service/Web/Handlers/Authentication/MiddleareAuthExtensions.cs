using System;
using Core.Services;
using Microsoft.AspNetCore.Authentication;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using Microsoft.Extensions.Options;

namespace Web.Handlers.Authentication
{
	public static class MiddlewareAuthExtensions
	{
		public static AuthenticationBuilder AddMiddlewareAuth<TService>(this AuthenticationBuilder builder) where TService : class, IAuthService
			=> builder.AddMiddlewareAuth<TService>(MiddlewareAuthDefaults.AuthenticationScheme, _ => { });

		public static AuthenticationBuilder AddMiddlewareAuth<TService>(this AuthenticationBuilder builder,
																		Action<MiddlewareAuthOptions> configureOptions) where TService : class, IAuthService
			=> builder.AddMiddlewareAuth<TService>(MiddlewareAuthDefaults.AuthenticationScheme, configureOptions);

		public static AuthenticationBuilder AddMiddlewareAuth<TService>(this AuthenticationBuilder builder, string authenticationScheme,
																		Action<MiddlewareAuthOptions> configureOptions) where TService : class, IAuthService
			=> builder.AddMiddlewareAuth<TService>(authenticationScheme, MiddlewareAuthDefaults.AuthenticationScheme, configureOptions);

		public static AuthenticationBuilder AddMiddlewareAuth<TService>(this AuthenticationBuilder builder, string authenticationScheme,
																		string displayName, Action<MiddlewareAuthOptions> configureOptions) where TService : class, IAuthService
		{
			builder.Services.AddSingleton<IAuthService, TService>();
			builder.Services.TryAddEnumerable(ServiceDescriptor.Singleton<IPostConfigureOptions<MiddlewareAuthOptions>, MiddlewareAuthPostConfigureOptions>());
			return builder.AddScheme<MiddlewareAuthOptions, MiddlewareAuthHandler>(authenticationScheme, displayName, configureOptions);
		}
	}
}