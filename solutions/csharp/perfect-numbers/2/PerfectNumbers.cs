using System;
using System.Collections.Generic;
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

        if (aliquotSum < number) 
            return Classification.Deficient;
        else if (aliquotSum > number) 
            return Classification.Abundant;
        else 
            return Classification.Perfect;
    }

    private static int GetAliquotSum(int number) => Enumerable.Range(1, number/2).Where(factor => number % factor == 0).Sum();
}
