using System;

public static class Raindrops
{
    public static string Convert(int number)
    {
        var rain = "";

        rain += number % 3 == 0 ? "Pling" : "";
        rain += number % 5 == 0 ? "Plang" : "";
        rain += number % 7 == 0 ? "Plong" : ""; 
        rain += rain.Equals("") ? number.ToString() : "";

        return rain;
    }
}