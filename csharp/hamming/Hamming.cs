using System;
using System.Linq;

public static class Hamming
{
    public static int Distance(string firstStrand, string secondStrand) => isEqualLength(firstStrand, secondStrand) ? _Distance(firstStrand, secondStrand) : throw new ArgumentException();

    private static int _Distance(string firstStrand, string secondStrand) => firstStrand.Zip(secondStrand, (firstChar, secondChar) => firstChar == secondChar).Where(x => x == false).Count();

    private static bool isEqualLength(string firstStrand, string secondStrand) => firstStrand.Length == secondStrand.Length;
}