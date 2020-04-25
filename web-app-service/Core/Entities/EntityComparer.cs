using System;
using System.Collections.Generic;

namespace Core.Entities
{
	public class EntityComparer<TEntity> : IEqualityComparer<TEntity> where TEntity : IBaseEntity
	{
		public bool Equals(TEntity x, TEntity y)
		{
			if (ReferenceEquals(x, y)) return true;
			if (ReferenceEquals(x, null) || ReferenceEquals(y, null)) return false;
			return x.UniqueID == y.UniqueID;
		}

		public int GetHashCode(TEntity entity) => entity.UniqueID?.GetHashCode() ?? 0;
	}
}