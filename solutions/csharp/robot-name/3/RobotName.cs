using System;
using System.Collections.Generic;

public class Robot
{
    private static string alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    private static HashSet<string> names = new HashSet<string>();
    private string _name;
    public string Name { get => _name; set => _name = value; }

    public Robot()
    {
        CreateName();
    }

    public void CreateName()
    {
        var random = new Random();

        for (int i = 0; i < 2; i++) Name += alphabet[random.Next(alphabet.Length)];
        for (int i = 0; i < 3; i++) Name += random.Next(10);

        if (!names.Add(Name)) Reset();
    }

    public void Reset()
    {
        Name = String.Empty;
        CreateName();
    }
}