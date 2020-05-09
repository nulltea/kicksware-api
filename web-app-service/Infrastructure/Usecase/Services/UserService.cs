using System.Collections.Generic;
using System.Threading.Tasks;
using Core.Entities.Users;
using Core.Gateway;
using Core.Repositories;
using Core.Services;

namespace Infrastructure.Usecase
{
	public class UserService : IUserService
	{
		private IUserRepository _repository;

		public UserService(IUserRepository repository) => _repository = repository;

		#region CRUD sync

		public User FetchUnique(string username, RequestParams requestParams = default) =>_repository.GetUnique(username, requestParams);

		public List<User> Fetch(RequestParams requestParams = default) =>_repository.Get(requestParams);

		public List<User> Fetch(Dictionary<string, object> query, RequestParams requestParams = default) =>
			_repository.Get(query, requestParams);

		public List<User> Fetch(IEnumerable<string> usernames, RequestParams requestParams = default) => _repository.Get(usernames, requestParams);

		public List<User> Fetch(object query, RequestParams requestParams = default) => _repository.Get(query, requestParams);

		public User Store(User user, RequestParams requestParams = default) => _repository.Post(user, requestParams);

		public bool Modify(User user, RequestParams requestParams = default) => _repository.Update(user, requestParams);

		public bool Remove(User user, RequestParams requestParams = default) => _repository.Delete(user, requestParams);

		public bool Remove(string username, RequestParams requestParams = default) => _repository.Delete(username, requestParams);

		public int Count(Dictionary<string, object> query, RequestParams requestParams = default) => _repository.Count(query, requestParams);

		public int Count(object query = default, RequestParams requestParams = default) => _repository.Count(query, requestParams);

		#endregion

		#region CRUD async

		public Task<User> FetchUniqueAsync(string username, RequestParams requestParams = default) => _repository.GetUniqueAsync(username, requestParams);

		public Task<List<User>> FetchAsync(RequestParams requestParams = default) => _repository.GetAsync(requestParams);

		public Task<List<User>> FetchAsync(Dictionary<string, object> query, RequestParams requestParams = default) => _repository.GetAsync(query, requestParams);

		public Task<List<User>> FetchAsync(IEnumerable<string> usernames, RequestParams requestParams = default) => _repository.GetAsync(usernames, requestParams);

		public Task<List<User>> FetchAsync(object query, RequestParams requestParams = default) => _repository.GetAsync(query, requestParams);

		public Task<User> StoreAsync(User user, RequestParams requestParams = default) => _repository.PostAsync(user, requestParams);

		public Task<bool> ModifyAsync(User user, RequestParams requestParams = default) => _repository.UpdateAsync(user, requestParams);

		public Task<bool> RemoveAsync(User user, RequestParams requestParams = default) => _repository.DeleteAsync(user, requestParams);

		public Task<bool> RemoveAsync(string username, RequestParams requestParams = default) => _repository.DeleteAsync(username, requestParams);

		public Task<int> CountAsync(Dictionary<string, object> query, RequestParams requestParams = default) => _repository.CountAsync(query, requestParams);

		public Task<int> CountAsync(object query = default, RequestParams requestParams = default) => _repository.CountAsync(requestParams);


		#endregion

		#region Usecase

		public void SendEmailConfirmation(string username, string callbackUrl) => throw new System.NotImplementedException();

		public Task SendEmailConfirmationAsync(string username, string callbackUrl) => throw new System.NotImplementedException();

		public void SendResetPasswordEmail(string username, string callbackUrl) => throw new System.NotImplementedException();

		public Task SendResetPasswordEmailAsync(string username, string callbackUrl) => throw new System.NotImplementedException();

		#endregion
	}
}