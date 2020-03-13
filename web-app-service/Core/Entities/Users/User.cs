﻿namespace Core.Entities.Users
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