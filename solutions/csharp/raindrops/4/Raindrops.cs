using System;
using System.Linq;
using System.Text;
using System.Collections.Generic;

public static class Raindrops
{
    private static readonly Dictionary<int, string> _sounds = new Dictionary<int, string>
    {
        { 3, "Pling" },
        { 5, "Plang" },
        { 7, "Plong" }
    };

    // new version with LINQ
    public static string Convert_LINQ(int number)
    {
        return string.Concat(_sounds
                             .Where(sound => number % sound.Key == 0)
                             .Select(sound => sound.Value)
                             .DefaultIfEmpty(number.ToString())
                             );
    }

    // new version of old iteration
    public static string Convert(int number)
    {
        var sound = new StringBuilder();
        if (number % 3 == 0) sound.Append("Pling");
        if (number % 5 == 0) sound.Append("Plang");
        if (number % 7 == 0) sound.Append("Plong");
        return (sound.Length == 0) ? number.ToString() : sound.ToString();
    }
}