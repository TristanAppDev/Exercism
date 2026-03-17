public static class Pangram
{
    private const string LETTERS = "abcdefghijklmnopqrstuvwxyz";
    public static bool IsPangram(string input)
    {
        string lowerCaseInput = input.ToLower();
        return LETTERS.All(lowerCaseInput.Contains);
    }
}
