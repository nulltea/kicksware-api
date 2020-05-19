using System.Threading.Tasks;
using Core.Entities;

namespace Core.Services.Interactive
{
	public interface ILikeService
	{
		public void Like(string entityID);

		public void Unlike(string entityID);

		public Task LikeAsync(string entityID);

		public Task UnlikeAsync(string entityID);
	}
}