using System.IO;

namespace Core.Constants
{
	public static class Constants
	{
		public const string GatewayBaseUrl = "http://localhost:8420";

		public static readonly string WebRootPath = Path.Combine(Directory.GetCurrentDirectory(), "wwwroot");

		public static readonly string FileStoragePath = Path.Combine(Directory.GetCurrentDirectory(), "wwwroot", "files");

	}
}