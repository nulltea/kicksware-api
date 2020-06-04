using System;
using System.IO;
using System.Text;

namespace Core.Constants
{
	public static class Constants
	{
		public static readonly string GatewayBaseUrl = Environment.GetEnvironmentVariable("GATEWAY_API_URL");

		public static readonly string FileStoragePath = Environment.GetEnvironmentVariable("STORAGE_PATH");

		public static readonly string WebRootPath = Path.Combine(Directory.GetCurrentDirectory(), "wwwroot");

		public static readonly string ImagesPath = Path.Combine(WebRootPath, "images");
	}
}
