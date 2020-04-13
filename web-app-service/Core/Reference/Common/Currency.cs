using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;

namespace Core.Reference
{
	public enum Currency
	{
		[EnumMember(Value = "USD")]
		[Display(Name = "U.S.Dollar", ShortName = "$")]
		UsDollar,

		[EnumMember(Value = "EUR")]
		[Display(Name = "Euro", ShortName = "€")]
		Euro,

		[EnumMember(Value = "GBP")]
		[Display(Name = "Pound Sterling", ShortName = "£")]
		PoundSterling,

		[EnumMember(Value = "UAH")]
		[Display(Name = "Ukrainian Hryvnia", ShortName = "₴")]
		UkrHryvnia,

		[EnumMember(Value = "RUB")]
		[Display(Name = "Russian Ruble", ShortName = "₽")]
		RusRuble,

		[EnumMember(Value = "PLN")]
		[Display(Name = "Poland Zloty", ShortName = "zł")]
		PolandZloty,

		[EnumMember(Value = "JPY")]
		[Display(Name = "Japanese Yen", ShortName = "¥")]
		JapanYen,

		[EnumMember(Value = "KRW")]
		[Display(Name = "South Korean won", ShortName = "₩")]
		KoreaWon
	}

	//<option value = "USD" > U.S.Dollar </ option >
	//< option value="EUR">Euro</option>
	//<option value = "GBP" > Pound Sterling</option>
	//<option value = "UAH" > Ukrainian Hryvnia</option>
	//<option value = "RUB" > Russian Ruble</option>
	//<option value = "PLN" > Poland Zloty</option>
	//<option value = "CAD" > Canadian Dollar</option>
	//<option value = "HKD" > Hong-Kong Dollar</option>
	//<option value = "JPY" > Japanese Yen</option>
	//<option value = "SGD" > Singapore Dollar</option>
	//<option value = "AUD" > Australian Dollar</option>
}