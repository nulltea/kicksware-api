using System.ComponentModel.DataAnnotations;

namespace Web.Models.Auth
{
	public class RecoveryCodeViewModel
	{
		[Required]
		[DataType(DataType.Text)]
		[Display(Name = "Recovery Code")]
		public string RecoveryCode { get; set; }
	}
}