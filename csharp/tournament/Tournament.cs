using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

public static class Tournament
{   
    public static void Tally(Stream inStream, Stream outStream)
    {



           var input = new StreamReader(inStream).ReadToEnd();
            var lines = input.Split("\n");
            var splittedLines = new List<string>();
            lines.ToList().ForEach(line => splittedLines.Add(line));

           Console.WriteLine(splittedLines);
           Console.WriteLine("===================================");
    }
}

class Team 
{
    private string name;
    private int matchesPlayed;
    private int wins;
    private int draws;
    private int loss;
    private int points;

    public Team(string name, int matchesPlayed, int wins, int draws, int loss, int points)
    {
        this.name = name;
        this.matchesPlayed = matchesPlayed;
        this.wins = wins;
        this.draws = draws;
        this.loss = loss;
        this.points = points;
    }
    public string Name { get => name; set => name = value; }
    public int MatchesPlayed { get => matchesPlayed; set => matchesPlayed = value; }
    public int Wins { get => wins; set => wins = value; }
    public int Draws { get => draws; set => draws = value; }
    public int Loss { get => loss; set => loss = value; }
    public int Points { get => points; set => points = value; }
} 