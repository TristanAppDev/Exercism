using System;
using System.Linq;

public static class Bob
{
    public static string Response(string message) => message switch
    {
        _ when message.IsSilence() => "Fine. Be that way!",
        _ when message.IsShoutingQuestion() => "Calm down, I know what I'm doing!",
        _ when message.IsShouting() => "Whoa, chill out!",
        _ when message.IsQuestion() => "Sure.",
        _ => "Whatever."
    };

    private static bool IsQuestion(this string message)
        => message.TrimEnd().EndsWith("?");

    private static bool IsShouting(this string message)
        => message.Any(char.IsLetter) && message.ToUpperInvariant() == message;

    private static bool IsSilence(this string message)
        => string.IsNullOrWhiteSpace(message);

    private static bool IsShoutingQuestion(this string message)
        => IsQuestion(message) && IsShouting(message);
}