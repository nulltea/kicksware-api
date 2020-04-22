using System.IO;
using System.Threading.Tasks;
using Core.Constants;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewEngines;
using Microsoft.AspNetCore.Mvc.ViewFeatures;

namespace Web.Utils.Extensions
{
	public static class ControllerExtensions
	{
		public static async Task<string> RenderViewAsync<TModel>(this Controller controller, string viewName, TModel model, bool isPartial = false)
		{
			if (string.IsNullOrEmpty(viewName)) viewName = controller.ControllerContext.ActionDescriptor.ActionName;

			controller.ViewData.Model = model;

			await using var writer = new StringWriter();
			var viewEngine = controller.HttpContext.RequestServices
				.GetService(typeof(ICompositeViewEngine)) as ICompositeViewEngine;
			var viewResult = GetViewEngineResult(controller, viewName, isPartial, viewEngine);

			if (!viewResult.Success) throw
				new System.Exception($"A view with the name {viewName} could not be found");

			var viewContext = new ViewContext(controller.ControllerContext, viewResult.View,
				controller.ViewData, controller.TempData, writer, new HtmlHelperOptions());

			await viewResult.View.RenderAsync(viewContext);

			return writer.GetStringBuilder().ToString();
		}

		private static ViewEngineResult GetViewEngineResult(ControllerBase controller, string viewName, bool isPartial, IViewEngine viewEngine)
		{
			if (!viewName.StartsWith("~/")) return viewEngine.FindView(controller.ControllerContext, viewName, !isPartial);
			var hostingEnv = controller.HttpContext.RequestServices.GetService(typeof(IWebHostEnvironment)) as
					IWebHostEnvironment;
			return viewEngine.GetView(hostingEnv?.WebRootPath ?? Constants.WebRootPath, viewName, !isPartial);
		}
	}
}