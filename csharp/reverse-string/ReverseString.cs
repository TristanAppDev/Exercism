using System;
using System.Linq;

public static class ReverseString
{
    public static string Reverse(string input)
    {
        return string.Join(string.Empty, input.Reverse().ToArray());
    }
}