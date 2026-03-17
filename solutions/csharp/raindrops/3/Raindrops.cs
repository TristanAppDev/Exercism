using System;

public static class Raindrops
{
    public static string Convert(int number)
    {
        var rain = string.Empty;
        if (number % 3 == 0) rain += "Pling";
        if (number % 5 == 0) rain += "Plang";
        if (number % 7 == 0) rain +=  "Plong"; 
        if (string.IsNullOrEmpty(rain)) rain = number.ToString();
        return rain;
    }
}