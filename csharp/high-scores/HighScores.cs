using System.Collections.Generic;
using System.Linq;

public class HighScores
{
    private readonly List<int> _scores;

    public HighScores(List<int> scores) => this._scores = scores;

    public List<int> Scores() => _scores;

    public int Latest() => _scores.Last();

    public int PersonalBest() => _scores.Max();

    public List<int> PersonalTopThree() => _scores
        .OrderByDescending(i => i)
        .Take(3)
        .ToList();
}