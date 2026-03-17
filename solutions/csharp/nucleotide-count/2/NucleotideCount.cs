using System;
using System.Linq;
using System.Collections.Generic;

public static class NucleotideCount
{
    public static IDictionary<char, int> Count(string sequence)
    {
        Dictionary<char, int> nucleotides = new Dictionary<char, int>
        {
            ['A'] = 0,
            ['C'] = 0,
            ['G'] = 0,
            ['T'] = 0
        };
        sequence.ToList()
                .ForEach(c => { if (nucleotides.ContainsKey(c)) { nucleotides[c] += 1; } else { throw new ArgumentException(); } });
        return nucleotides;
    }
}