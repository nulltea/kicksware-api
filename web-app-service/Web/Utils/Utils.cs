using System;
using System.Collections.Generic;
using System.Linq;
using Core.Entities.Products;
using Core.Entities.Users;
using Web.Models;

namespace Web.Utils
{
	public static class Utils
	{
		public static List<SneakerProductViewModel> ToViewModel(this List<SneakerProduct> entities) =>
			entities.CastExtend<SneakerProduct, SneakerProductViewModel>();

		public static List<SneakerProduct> ToEntityModel(this List<SneakerProductViewModel> viewModels) =>
			viewModels.Cast<SneakerProduct>().ToList();

		public static List<UserViewModel> ToViewModel(this List<User> entities) =>
			entities.CastExtend<User, UserViewModel>();

		public static List<User> ToEntityModel(this List<UserViewModel> viewModels) => viewModels.Cast<User>().ToList();

		public static SneakerProductViewModel ToViewModel(this SneakerProduct entity) =>
			entity.CastExtend<SneakerProduct, SneakerProductViewModel>();

		public static UserViewModel ToViewModel(this User entity) => entity.CastExtend<User, UserViewModel>();

		public static TTarget CastExtend<TSource, TTarget>(this TSource entity) where TTarget : TSource, new()
		{
			var instance = Activator.CreateInstance<TTarget>();
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

		public static List<TTarget> CastExtend<TSource, TTarget>(this List<TSource> entities) where TTarget : TSource, new()
		{
			return entities.Select(CastExtend<TSource, TTarget>).ToList();
		}
	}

}