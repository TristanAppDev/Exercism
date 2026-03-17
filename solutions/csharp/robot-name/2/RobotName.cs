using System;
using System.Collections.Generic;

public class Robot
{

    private static HashSet<string> names = new HashSet<string>();

    private string _name;
    public string Name
    {
        get
        {
            if (string.IsNullOrEmpty(Name)) CreateName();
            return Name;
        }
        private set => _name = value;
    }

    public Robot() => CreateName();

    public void CreateName()
    {
        var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        var random = new Random();

        for (int i = 0; i < 2; i++) Name += alphabet[random.Next(alphabet.Length)];
        for (int i = 0; i < 3; i++) Name += random.Next(10);

        if (!names.Add(Name)) {
            Reset();
            CreateName();
        }
    }

    public void Reset()
    {
        Name = String.Empty;
    }
}