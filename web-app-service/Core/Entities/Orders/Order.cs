using System;
using Core.Attributes;

namespace Core.Entities
{
	[EntityService(Resource = "orders")]
	public class Order
	{
		public string UserID { get; set; }

		public string EntityID { get; set; }

		public DateTime OrderedAt { get; set; }
	}
}