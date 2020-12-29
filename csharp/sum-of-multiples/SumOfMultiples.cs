using System.Linq;
using System.Collections.Generic;

public static class SumOfMultiples
{
    public static int Sum(IEnumerable<int> multiples, int max) => getMultiples(multiples, max).Sum();

    private static IEnumerable<int> getMultiples(IEnumerable<int> multiples, int max) => Enumerable.Range(1, max - 1).Where(num => multiples.Any(factor => factor > 0 && num % factor == 0));
}