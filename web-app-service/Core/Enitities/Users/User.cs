using System.Collections.Generic;
using Core.Enitities.Products;

namespace Core.Enitities.Users
{
	public class User : IBaseEntity
	{
		public string UniqueId { get; set; }

		public string UserName { get; set; }

		public string PasswordHash { get; set; }

		public string FirstName { get; set; }

		public string LastName { get; set; }
	}
}