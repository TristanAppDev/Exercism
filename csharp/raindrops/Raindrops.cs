using System;
using System.Linq;
using System.Text;

public static class Raindrops
{
    private static readonly Dictionary<int, string> _sounds = new Dictionary<int, string>
    {
        { 3, "Pling" },
        { 5, "Pling" },
        { 7, "Pling" }
    };

    public static string Convert(int number)
    {
        return string
            .Concat(_sounds
            .Where(soundVal => number % soundVal.Key == 0)
            .Select(soundVal => soundVal.Value)
            .DefaultIfEmpty(number.ToString())
            );
    }

    public static string Convert_(int number)
    {
        var sound = new StringBuilder();
        if (number % 3 == 0) sound.Append("Pling");
        if (number % 5 == 0) sound.Append("Plang");
        if (number % 7 == 0) sound.Append("Plong");
        return (sound.Length == 0) ? number.ToString() : sound.ToString();
    }
}