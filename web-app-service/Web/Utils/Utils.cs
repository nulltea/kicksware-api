using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using Core.Entities.Products;
using Core.Entities.Users;
using web_app_service.Models;

namespace web_app_service.Utils
{
	public static class Utils
	{
		public static List<SneakerProductViewModel> ToViewModel(this List<SneakerProduct> entities) =>
			entities.Select(CastExtend<SneakerProductViewModel>).ToList();

		public static List<SneakerProduct> ToEntityModel(this List<SneakerProductViewModel> viewModels) =>
			viewModels.Cast<SneakerProduct>().ToList();

		public static List<UserViewModel> ToViewModel(this List<User> entities) => entities.Select(CastExtend<UserViewModel>).ToList();

		public static List<User> ToEntityModel(this List<UserViewModel> viewModels) => viewModels.Cast<User>().ToList();

		public static SneakerProductViewModel ToViewModel(this SneakerProduct entity) => CastExtend<SneakerProductViewModel>(entity);

		public static UserViewModel ToViewModel(this User entity) => CastExtend<UserViewModel>(entity);

		private static T CastExtend<T>(object entity)
		{
			var instance = Activator.CreateInstance<T>();
			var type = entity.GetType();
			var properties = type.GetProperties();
			foreach (var property in properties)
			{
				if (property.CanWrite)
				{
					property.SetValue(instance, property.GetValue(entity, null), null);
				}
			}
			return instance;
		}
	}
}