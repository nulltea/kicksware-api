using System.Runtime.Serialization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace Core.Reference
{
	public enum SortDirection
	{
		[EnumMember(Value = "asc")]
		Ascending,

		[EnumMember(Value = "desc")]
		Descending
	}
}