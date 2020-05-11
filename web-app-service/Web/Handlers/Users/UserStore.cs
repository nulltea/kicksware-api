using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Identity;
using Core.Entities.Users;
using Core.Services;

namespace Web.Handlers.Users
{
	public class UserStore : IUserEmailStore<User>, IUserPasswordStore<User>, IUserPhoneNumberStore<User>
	{
		private IUserService _service;

		public UserStore(IUserService service) => _service = service;

		public async Task<IdentityResult> CreateAsync(User user, CancellationToken _)
		{
			await _service.StoreAsync(user);
			return IdentityResult.Success;
		}

		public async Task<IdentityResult> DeleteAsync(User user, CancellationToken _)
		{
			await _service.RemoveAsync(user);
			return IdentityResult.Success;
		}

		public Task<User> FindByIdAsync(string userId, CancellationToken _) => _service.FetchUniqueAsync(userId);

		public Task<User> FindByNameAsync(string normalizedUserName, CancellationToken _) => _service.FetchUniqueAsync(normalizedUserName.ToLower());

		public Task<string> GetNormalizedUserNameAsync(User user, CancellationToken _) => Task.FromResult(user.Email.Split("@")[0]);

		public Task<string> GetUserIdAsync(User user, CancellationToken _) => Task.FromResult(user.UniqueID);

		public Task<string> GetUserNameAsync(User user, CancellationToken _) => Task.FromResult(user.Email.Split("@")[0]);

		public Task SetNormalizedUserNameAsync(User user, string normalizedName, CancellationToken _)
		{
			user.UniqueID = normalizedName;
			return Task.CompletedTask;
		}

		public Task SetUserNameAsync(User user, string userName, CancellationToken _)
		{
			user.UniqueID = userName;
			return Task.CompletedTask;
		}

		public async Task<IdentityResult> UpdateAsync(User user, CancellationToken _)
		{
			await _service.ModifyAsync(user);
			return IdentityResult.Success;
		}

		public async Task<User> FindByEmailAsync(string normalizedEmail, CancellationToken _) =>
			(await _service.FetchAsync(PropertyQuery("email", normalizedEmail))).FirstOrDefault();

		public Task<string> GetEmailAsync(User user, CancellationToken _) => Task.FromResult(user.Email);

		public Task<bool> GetEmailConfirmedAsync(User user, CancellationToken _) => Task.FromResult(user.Confirmed);

		public Task<string> GetNormalizedEmailAsync(User user, CancellationToken _) => Task.FromResult(user.Email);

		public Task SetEmailAsync(User user, string email, CancellationToken _)
		{
			user.Email = email;
			return Task.CompletedTask;
		}

		public Task SetEmailConfirmedAsync(User user, bool confirmed, CancellationToken _)
		{
			user.Confirmed = confirmed;
			return Task.CompletedTask;
		}

		public Task SetNormalizedEmailAsync(User user, string normalizedEmail, CancellationToken _)
		{
			user.Email = normalizedEmail;
			return Task.CompletedTask;
		}

		public Task<string> GetPasswordHashAsync(User user, CancellationToken _) => Task.FromResult(user.PasswordHash);

		public Task<bool> HasPasswordAsync(User user, CancellationToken _) => Task.FromResult(string.IsNullOrEmpty(user.PasswordHash));

		public Task SetPasswordHashAsync(User user, string passwordHash, CancellationToken _)
		{
			user.PasswordHash = passwordHash;
			return Task.CompletedTask;
		}

		public Task<string> GetPhoneNumberAsync(User user, CancellationToken _) => Task.FromResult(user.PhoneNumber);

		public Task<bool> GetPhoneNumberConfirmedAsync(User user, CancellationToken _) => Task.FromResult(true);

		public Task SetPhoneNumberAsync(User user, string phoneNumber, CancellationToken _)
		{
			user.PhoneNumber = phoneNumber;
			return Task.CompletedTask;
		}

		public Task SetPhoneNumberConfirmedAsync(User user, bool confirmed, CancellationToken _) => Task.CompletedTask;


		public void Dispose() { }

		public Dictionary<string, object> PropertyQuery(string prop, object value) =>
			new Dictionary<string, object> {{prop, value}};
	}
}