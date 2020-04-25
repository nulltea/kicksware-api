using System;
using System.Collections.Generic;
using System.Linq;
using Core.Algorithms;

namespace Core.Extension
{
	public static class SequenceExtension
	{
		public static IOrderedEnumerable<TSource> OrderBySimilarity<TSource>(this IEnumerable<TSource> sequence, Func<TSource, string> selector, string pattern)
		{
			return sequence.OrderBy(item => LevenshteinDistance.Calculate(selector(item), pattern));
		}
	}
}