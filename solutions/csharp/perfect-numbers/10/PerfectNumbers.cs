using System;
using System.Linq;

public enum Classification
{
    Perfect,
    Abundant,
    Deficient
}

public static class PerfectNumbers
{
    public static Classification Classify(int number)
    {
        if (number <= 0) throw new ArgumentOutOfRangeException();
        return (number - GetAliquotSum(number)) switch
        {
            > 0 => Classification.Deficient,
            < 0 => Classification.Abundant,
            _ => Classification.Perfect,
        };
    }

    private static int GetAliquotSum(int number) =>
        Enumerable
        .Range(1, number / 2)
        .Where(factor => number % factor == 0)
        .Sum();
}