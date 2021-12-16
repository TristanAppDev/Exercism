using System;
using System.Linq;
using System.Text;

public static class RotationalCipher
{
    public static string Rotate(string text, int key) => new StringBuilder().AppendJoin("", text.Select(c => Shift(c, key))).ToString();
    
    private static char Shift(char c, int key) => Char.IsLetter(c) ? ShiftLetter(c, key) : c;
    
    private static char ShiftLetter(char letter, int key)
    {
        var offset = Char.IsLower(letter) ? 'a' : 'A';
        return (char)((letter + key - offset) % 26 + offset);
    }
}
