using System;
using System.Linq;

public static class DifferenceOfSquares
{
    public static int Pow(int number) => number * number;
    
    public static int CalculateSquareOfSum(int n) => Pow((n * (n + 1)) / 2);

    public static int CalculateSumOfSquares(int n) => (n * (n + 1) * (2 * n + 1)) / 6;

    public static int CalculateDifferenceOfSquares(int max) => 
        CalculateSquareOfSum(max) - CalculateSumOfSquares(max);
}