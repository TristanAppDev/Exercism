using System;

public class Player
{
    private const int DIE_SIGHTS = 18;
    private Random rnd;

    public Player() => rnd = new Random();

    public int RollDie() => rnd.Next(1, DIE_SIGHTS + 1);

    public double GenerateSpellStrength() => rnd.NextDouble() * 100;
}
