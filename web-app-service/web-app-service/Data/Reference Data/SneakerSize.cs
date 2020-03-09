namespace web_app_service.Data.Reference_Data
{
	public struct SneakerSize
	{
		public decimal Europe { get; set; }
		public decimal UnitedStates { get; set; }
		public decimal UnitedKingdom { get; set; }
		public decimal Centimeters { get; set; }
		public override string ToString() => $"{Europe} EU | {UnitedStates} US | {UnitedKingdom} UK";
	}

	public static partial class ReferenceData
	{
		public static readonly SneakerSize[] SneakerSizesList =
		{
			new SneakerSize
			{
				Europe = 38m,
				UnitedStates = 6m,
				UnitedKingdom = 5m,
				Centimeters = 24
			},
			new SneakerSize
			{
				Europe = 39m,
				UnitedStates = 6.5m,
				UnitedKingdom = 5.5m,
				Centimeters = 24.5m
			},
			new SneakerSize
			{
				Europe = 39.5m,
				UnitedStates = 7m,
				UnitedKingdom = 6m,
				Centimeters = 25
			},
			new SneakerSize
			{
				Europe = 40m,
				UnitedStates = 7.5m,
				UnitedKingdom = 6.5m,
				Centimeters = 25.5m
			},
			new SneakerSize
			{
				Europe = 41m,
				UnitedStates = 8m,
				UnitedKingdom = 7m,
				Centimeters = 26
			},
			new SneakerSize
			{
				Europe = 41.5m,
				UnitedStates = 8.5m,
				UnitedKingdom = 7.5m,
				Centimeters = 26.5m
			},
			new SneakerSize
			{
				Europe = 42m,
				UnitedStates = 9m,
				UnitedKingdom = 8m,
				Centimeters = 27
			},
			new SneakerSize
			{
				Europe = 42.5m,
				UnitedStates = 9.5m,
				UnitedKingdom = 8.5m,
				Centimeters = 27.5m
			},
			new SneakerSize
			{
				Europe = 43m,
				UnitedStates = 10m,
				UnitedKingdom = 9m,
				Centimeters = 28
			},
			new SneakerSize
			{
				Europe = 44m,
				UnitedStates = 10.5m,
				UnitedKingdom = 9.5m,
				Centimeters = 28.5m
			},
			new SneakerSize
			{
				Europe = 44.5m,
				UnitedStates = 11m,
				UnitedKingdom = 10m,
				Centimeters = 29
			},
			new SneakerSize
			{
				Europe = 45m,
				UnitedStates = 11.5m,
				UnitedKingdom = 10.5m,
				Centimeters = 29.5m
			},
			new SneakerSize
			{
				Europe = 46m,
				UnitedStates = 12m,
				UnitedKingdom = 11m,
				Centimeters = 30
			},
			new SneakerSize
			{
				Europe = 46.5m,
				UnitedStates = 12.5m,
				UnitedKingdom = 11.5m,
				Centimeters = 30.5m
			},
			new SneakerSize
			{
				Europe = 47m,
				UnitedStates = 13m,
				UnitedKingdom = 12m,
				Centimeters = 31
			},
		};
	}

	/*
	   <option value="5">US 5 / EU 37</option>
                                            <option value="5.5">US 5.5 / EU 38</option>
                                            <option value="6">US 6 / EU 39</option>
                                            <option value="6.5">US 6.5 / EU 39-40</option>
                                            <option value="7">US 7 / EU 40</option>
                                            <option value="7.5">US 7.5 / EU 40-41</option>
                                            <option value="8">US 8 / EU 41</option>
                                            <option value="8.5">US 8.5 / EU 41-42</option>
                                            <option value="9">US 9 / EU 42</option>
                                            <option value="9.5">US 9.5 / EU 42-43</option>
                                            <option value="10">US 10 / EU 43</option>
                                            <option value="10.5">US 10.5 / EU 43-44</option>
                                            <option value="11">US 11 / EU 44</option>
                                            <option value="11.5">US 11.5 / EU 44-45</option>
                                            <option value="12">US 12 / EU 45</option>
                                            <option value="12.5">US 12.5 / EU 45-46</option>
                                            <option value="13">US 13 / EU 46</option>
                                            <option value="14">US 14 / EU 47</option>
                                            <option value="15">US 15 / EU 48</option>
	 
	*/
	
}
 