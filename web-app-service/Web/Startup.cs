using System;
using Core.Entities.Products;
using Core.Entities.References;
using Core.Gateway;
using Core.Model;
using Core.Repositories;
using Core.Services;
using Infrastructure.Data;
using Infrastructure.Gateway.REST;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Usecase;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.StaticFiles;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using Microsoft.Extensions.Hosting;
using Microsoft.IdentityModel.Tokens;
using SmartBreadcrumbs.Extensions;
using Web.Data;
using Web.Handlers.Filter;

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
			services.AddDefaultIdentity<IdentityUser>(options => options.SignIn.RequireConfirmedAccount = true);
			services.AddControllersWithViews();
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

			services.AddSingleton<IGatewayClient<IGatewayRestRequest>, RestfulClient>();

			services.AddSingleton<ISneakerProductRepository, SneakerProductsRestRepository>();
			services.AddSingleton<ISneakerReferenceRepository, SneakerReferencesRestRepository>();

			services.AddTransient<ICommonService<SneakerReference>, SneakerReferenceService>();
			services.AddTransient<ICommonService<SneakerProduct>, SneakerProductService>();
			services.AddSingleton<ISneakerReferenceService, SneakerReferenceService>();
			services.AddSingleton<ISneakerProductService, SneakerProductService>();
			services.AddSingleton<IReferenceSearchService, ReferenceSearchService>();

			services.AddTransient<FilterContentBuilder<SneakerReference>, ReferencesFilterContent>();
			services.AddTransient<FilterContentBuilder<SneakerProduct>, ProductsFilterContent>();

			#endregion

			#region Authentication

			services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
				.AddCookie()
				.AddJwtBearer(options => ConfigureJwtAuth(options))
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

		private JwtBearerOptions ConfigureJwtAuth(JwtBearerOptions options)
		{
			options.RequireHttpsMetadata = false;
			options.TokenValidationParameters = new TokenValidationParameters
			{
				SaveSigninToken = true,
			};
			return options;
		}

		#endregion
	}
}
