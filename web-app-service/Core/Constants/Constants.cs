using System.IO;

namespace Core.Constants
{
	public static class Constants
	{
		public const string GatewayBaseUrl = "https://api.kicksware.com";

		public static readonly string WebRootPath = Path.Combine(Directory.GetCurrentDirectory(), "wwwroot");

		public static readonly string FileStoragePath = Path.Combine(WebRootPath, "files");

		public static readonly string ImagesPath = Path.Combine(WebRootPath, "images");
	}
}
