using System.Text;
using Core.Entities.Users;
using Microsoft.AspNetCore.Authentication;
using Newtonsoft.Json;

namespace Web.Handlers.Authentication
{
	public class AuthTokenSerializer : IDataSerializer<AuthToken>
	{
		public AuthToken Deserialize(byte[] data) => JsonConvert.DeserializeObject<AuthToken>(Encoding.UTF8.GetString(data));

		public byte[] Serialize(AuthToken token) => Encoding.UTF8.GetBytes(JsonConvert.SerializeObject(token));
	}
}