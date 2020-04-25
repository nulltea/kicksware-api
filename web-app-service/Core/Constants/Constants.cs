using System.IO;

namespace Core.Constants
{
	public static class Constants
	{
		public const string GatewayBaseUrl = "http://kicksware.com:8080/api";

		public static readonly string WebRootPath = Path.Combine(Directory.GetCurrentDirectory(), "wwwroot");

		public static readonly string FileStoragePath = Path.Combine(WebRootPath, "files");

		public static readonly string ImagesPath = Path.Combine(WebRootPath, "images");
	}
}