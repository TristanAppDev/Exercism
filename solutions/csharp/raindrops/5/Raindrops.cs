using System;
using System.Linq;
using System.Text;
using System.Collections.Generic;

public static class Raindrops
{
    private static readonly List<(int Factor, string Sound)> _raindrops = new()
    {
        ( Factor: 3, Sound: "Pling" ),
        ( Factor: 5, Sound: "Plang" ),
        ( Factor: 7, Sound: "Plong" )
    };

    // new version with LINQ
    public static string Convert_LINQ(int number)
    {
        return string.Concat(_raindrops
                             .Where(raindrop => number % raindrop.Factor == 0)
                             .Select(raindrop => raindrop.Sound)
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