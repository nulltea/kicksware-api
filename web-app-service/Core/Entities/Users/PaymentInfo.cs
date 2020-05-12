using System;

namespace Core.Entities.Users
{
	public class PaymentInfo
	{
		public string CardNumber { get; set; }

		public DateTime Expires { get; set; }

		public string CVV { get; set; }

		public AddressInfo BillingInfo { get; set; } = new AddressInfo();
	}
}