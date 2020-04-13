namespace Core.Reference
{
	public struct SneakerSize
	{
		public decimal Europe { get; set; }
		public decimal UnitedStates { get; set; }
		public decimal UnitedKingdom { get; set; }
		public decimal Centimeters { get; set; }
		public override string ToString() => $"{Europe} EU | {UnitedStates} US | {UnitedKingdom} UK";
	}
}
 