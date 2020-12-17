using System;
using System.Collections.Generic;
using System.Linq;

public static class RnaTranscription
{
    public static string ToRna(string nucleotide)
    {
        var dnaRna = new Dictionary<char, char> { { 'G', 'C' }, { 'C', 'G' }, { 'T', 'A' }, { 'A', 'U' } };
        return string.Join("", nucleotide.Select(dna => dna = dnaRna.ContainsKey(dna) ? dnaRna[dna] : dna).ToList());
    }
}