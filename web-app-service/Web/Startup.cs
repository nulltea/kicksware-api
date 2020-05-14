using System;
using System.Collections.Generic;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Entities.Users;
using Core.Gateway;
using Core.Model;
using Core.Repositories;
using Core.Services;
using Infrastructure.Data;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Usecase;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Identity;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using SmartBreadcrumbs.Extensions;
using Web.Container.Factory;
using Web.Handlers.Authentication;
using Web.Handlers.Authorisation;
using Web.Handlers.Filter;
using Web.Handlers.Users;

namespace Web
{
	public class Startup
	{
		private IWebHostEnvironment HostEnvironment { get; }

		public Startup(IConfiguration configuration, IWebHostEnvironment env) => (Configuration, HostEnvironment) = (configuration, env);

		private IConfiguration Configuration { get; }

		// This method gets called by the runtime. Use this method to add services to the container.
		public void ConfigureServices(IServiceCollection services)
		{
			services.AddDefaultIdentity<User>(ConfigureAuthOptions);
			services.AddControllersWithViews();
			services.AddHttpContextAccessor();
			services.AddSession();

			services.AddBreadcrumbs(GetType().Assembly, options =>
			{
				options.TagName = "nav";
				options.TagClasses = "";
				options.OlClasses = "breadcrumb";
				options.LiClasses = "breadcrumb-item";
				options.ActiveLiClasses = "breadcrumb-item active";
				options.SeparatorElement = "<li class=\"separator\">\u276F</li>";
			});

			var builder = services.AddRazorPages();
#if DEBUG
			if (HostEnvironment.IsDevelopment()) builder.AddRazorRuntimeCompilation();
#endif

			#region Dependency injection

			services.AddTransient<IGatewayClient<IGatewayRestRequest>, RestfulClient>(ServiceFactory.ProvideRestClient);

			services.AddSingleton<ISneakerProductRepository, SneakerProductsRestRepository>();
			services.AddSingleton<ISneakerReferenceRepository, SneakerReferencesRestRepository>();
			services.AddSingleton<IUserRepository, UserRestRepository>();

			services.AddTransient<ICommonService<SneakerReference>, SneakerReferenceService>();
			services.AddTransient<ICommonService<SneakerProduct>, SneakerProductService>();
			services.AddTransient<ISneakerReferenceService, SneakerReferenceService>();
			services.AddTransient<ISneakerProductService, SneakerProductService>();
			services.AddTransient<IReferenceSearchService, ReferenceSearchService>();
			services.AddTransient<IUserService, UserService>();

			services.AddTransient<FilterContentBuilder<SneakerReference>, ReferencesFilterContent>();
			services.AddTransient<FilterContentBuilder<SneakerProduct>, ProductsFilterContent>();

			services.AddTransient<IUserStore<User>, UserStore>();
			services.AddTransient<SignInManager<User>, MiddlewareSignInManager>();
			services.AddTransient<IAuthorizationHandler, NotGuestHandler>();

			services.AddSingleton<ISecureDataFormat<AuthToken>, SecureDataFormat<AuthToken>>(ServiceFactory.ProvideSecureTokenFormat);

			#endregion

			#region Authentication

			services
				.AddAuthentication(ConfigureAuthOptions)
				.AddMiddlewareAuth<AuthService>(ConfigureAuthOptions)
				.AddFacebook(facebookOptions =>
				{
					facebookOptions.AppId = Environment.GetEnvironmentVariable("Authentication:Facebook:AppId");
					facebookOptions.AppSecret = Environment.GetEnvironmentVariable("Authentication:Facebook:AppSecret");
				})
				.AddGoogle(options =>
				{
					options.ClientId = Environment.GetEnvironmentVariable("Authentication:Google:ClientId");
					options.ClientSecret = Environment.GetEnvironmentVariable("Authentication:Google:ClientSecret");
				});

			services.AddAuthorization(ConfigureAuthOptions);

			#endregion
		}

		// This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
		public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
		{
			if (env.IsDevelopment())
			{
				app.UseDeveloperExceptionPage();
			}
			else
			{
				app.UseExceptionHandler("/Home/Error");
				app.UseHsts();
			}
			app.UseHttpsRedirection();
			app.UseStaticFiles();

			app.UseRouting();
			app.UseSession();

			app.UseAuthentication();
			app.UseAuthorization();

			app.UseEndpoints(endpoints =>
			{
				endpoints.MapControllerRoute(
					name: "default",
					pattern: "{controller=Home}/{action=Index}/{id?}"
				);
				endpoints.MapRazorPages();
			});
		}

		#region Configuration handlers

		private static void ConfigureAuthOptions(AuthenticationOptions options)
		{
			options.DefaultScheme = MiddlewareAuthDefaults.AuthenticationScheme;
			options.DefaultSignInScheme = MiddlewareAuthDefaults.AuthenticationScheme;
			options.DefaultAuthenticateScheme = MiddlewareAuthDefaults.AuthenticationScheme;
			options.DefaultSignOutScheme = MiddlewareAuthDefaults.AuthenticationScheme;
			options.DefaultChallengeScheme = MiddlewareAuthDefaults.AuthenticationScheme;
			options.SchemeMap[IdentityConstants.ApplicationScheme].HandlerType = typeof(MiddlewareAuthHandler);
		}

		private static void ConfigureAuthOptions(IdentityOptions options)
		{
			options.SignIn.RequireConfirmedEmail = true;
			options.SignIn.RequireConfirmedPhoneNumber = false;

			options.Password.RequireUppercase = true;
			options.Password.RequireLowercase = true;
			options.Password.RequireNonAlphanumeric = true;
			options.Password.RequireDigit = true;
			options.Password.RequiredLength = 6;

			options.User.RequireUniqueEmail = true;
		}

		private static void ConfigureAuthOptions(AuthorizationOptions options)
		{
			options.AddPolicy("NotGuest", policy => policy.Requirements.Add(new NotGuestRequirement()));
		}

		private static void ConfigureAuthOptions(MiddlewareAuthOptions options)
		{

		}

		#endregion
	}
}
