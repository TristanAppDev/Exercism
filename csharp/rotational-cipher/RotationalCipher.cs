using System;
using System.Linq;

public static class RotationalCipher
{
    public static string Rotate(string text, int shiftKey) => 
        string.Concat(text.Select(c => Char.IsLetter(c) ? shiftLetter(c, shiftKey, Char.IsUpper(c) ? 'A' : 'a') : c));
        
    public static char shiftLetter(char letter, int shiftKey, int lowerUpperConst) => 
        (char)((letter + shiftKey - lowerUpperConst) % 26 + lowerUpperConst);
}
