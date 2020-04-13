namespace Core.Reference
{
	public class ShippingInfo
	{
		public bool Possible { get; set; } = true;

		public decimal Cost { get; set; }

		public ShippingInfo() { }

		public ShippingInfo(bool possible, decimal cost) => (Possible, Cost) = (possible, cost);
	}
}