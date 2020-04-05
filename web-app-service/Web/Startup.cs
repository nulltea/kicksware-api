using System;
using Core.Repositories;
using Core.Services;
using Infrastructure.Data;
using Infrastructure.Gateway.REST.Client;
using Infrastructure.Usecase;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.StaticFiles;
using Microsoft.EntityFrameworkCore;
using web_app_service.Data;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

namespace web_app_service
{
	public class Startup
	{
		public IWebHostEnvironment HostEnvironment { get; set; }

		public Startup(IConfiguration configuration, IWebHostEnvironment env) => (Configuration, HostEnvironment) = (configuration, env);

		public IConfiguration Configuration { get; }

		// This method gets called by the runtime. Use this method to add services to the container.
		public void ConfigureServices(IServiceCollection services)
		{
			services.AddDbContext<ApplicationDbContext>(options =>
				options.UseSqlServer(
					Configuration.GetConnectionString("DefaultConnection")));
			services.AddDefaultIdentity<IdentityUser>(options => options.SignIn.RequireConfirmedAccount = true)
				.AddEntityFrameworkStores<ApplicationDbContext>();
			services.AddControllersWithViews();
			var builder = services.AddRazorPages();
#if DEBUG
			if (HostEnvironment.IsDevelopment()) builder.AddRazorRuntimeCompilation();
#endif

			#region Dependency injection

			services.AddSingleton<RestfulClient, RestfulClient>();
			services.AddSingleton<ISneakerProductRepository, SneakerProductsRestRepository>();
			services.AddSingleton<ISneakerProductService, SneakerProductService>();
			services.AddSingleton<IReferenceSearchService, ReferenceSearchService>();

			#endregion

			#region Authentication

			services.AddAuthentication()
				.AddFacebook(facebookOptions =>
				{
					facebookOptions.AppId = Environment.GetEnvironmentVariable("Authentication:Facebook:AppId"); //Configuration["Authentication:Facebook:AppId"];
					facebookOptions.AppSecret = Environment.GetEnvironmentVariable("Authentication:Facebook:AppSecret"); //Configuration["Authentication:Facebook:AppSecret"];
				})
				.AddGoogle(options =>
				{
					options.ClientId = Environment.GetEnvironmentVariable("Authentication:Google:ClientId");//Configuration["Authentication:Google:ClientId"];
					options.ClientSecret = Environment.GetEnvironmentVariable("Authentication:Google:ClientSecret"); //Configuration["Authentication:Google:ClientSecret"];
				});

			#endregion
		}

		// This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
		public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
		{
			if (env.IsDevelopment())
			{
				app.UseDeveloperExceptionPage();
				app.UseDatabaseErrorPage();
			}
			else
			{
				app.UseExceptionHandler("/Home/Error");
				// The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
				app.UseHsts();
			}
			app.UseHttpsRedirection();
			app.UseStaticFiles();

			app.UseRouting();

			app.UseAuthentication();
			app.UseAuthorization();

			app.UseEndpoints(endpoints =>
			{
				endpoints.MapControllerRoute(
					name: "default",
					pattern: "{controller=Home}/{action=Index}/{id?}");
				endpoints.MapRazorPages();
			});

			var provider = new FileExtensionContentTypeProvider();
			provider.Mappings[".less"] = "plain/text";

			app.UseStaticFiles(new StaticFileOptions
			{
				ContentTypeProvider = provider
			});
		}
	}
}
