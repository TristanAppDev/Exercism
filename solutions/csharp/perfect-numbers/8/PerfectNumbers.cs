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
        var aliquotSum = GetAliquotSum(number);
        return aliquotSum < number ? Classification.Deficient
            : aliquotSum > number ? Classification.Abundant
            : Classification.Perfect;
    }

    private static int GetAliquotSum(int number) => 
        Enumerable
        .Range(1, number / 2)
        .Where(factor => number % factor == 0)
        .Sum(); 
}