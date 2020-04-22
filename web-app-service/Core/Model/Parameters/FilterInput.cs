using Newtonsoft.Json;

namespace Core.Model.Parameters
{
	public class FilterInput
	{
		public string RenderId { get; set; }

		public bool Checked { get; set; }

		public object Value => JsonConvert.DeserializeObject(ValueJson ?? JsonConvert.Null);

		public string ValueJson { get; set; }
	}
}