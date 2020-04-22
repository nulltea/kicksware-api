using System.Collections.Generic;
using Core.Entities;
using Core.Extension;

namespace Core.Model
{
	public abstract class FilterContentBuilder<TEntity> where TEntity : IBaseEntity
	{
		protected Dictionary<string, object> AdditionalParams { get; set; }

		public virtual void SetAdditionalParams(object param) => AdditionalParams = param?.ToMap();
		public abstract void ConfigureFilter(IFilteredModel<TEntity> model);
		public abstract void ConfigureSorting(IFilteredModel<TEntity> model);
	}
}