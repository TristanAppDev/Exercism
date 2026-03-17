using System.Linq;

public static class Pangram
{
    public static bool IsPangram(string input)
    {
        return Enumerable.Range('a', 26)
        .Select(x => (char)x)
        .ToArray()
        .Count(c => input.ToLower().Contains(c)) == 26;
    }
}
