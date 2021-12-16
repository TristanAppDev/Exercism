using System;
using System.Collections.Generic;
using System.Linq;

public static class ReverseString
{

    static List<int> _Scores;
    public static string Reverse(string input)
    {
        return string.Join(string.Empty, input.Reverse().ToArray());
    }

    public static string Reverse_ugly(string input)
    {
        string aux="";
        for (int i = 0; i < input.Length; i++)
        {
           aux=input[i]+aux;
        }
        return aux;
    }

    public static void HighScores(List<int> scores) {
        List<int> allScores = new List<int>();
        var concat_scores = allScores.Concat(scores);
        _Scores = scores;
        
        Console.WriteLine(_Scores.GetHashCode());
        Console.WriteLine(concat_scores.GetHashCode());
        Console.WriteLine(scores.GetHashCode());

        _Scores = null;
    }
}