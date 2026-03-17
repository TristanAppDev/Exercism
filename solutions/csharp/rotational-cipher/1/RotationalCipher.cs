using System;
using System.Collections.Generic;
using System.Linq;

public static class RotationalCipher
{
    public static string Rotate(string text, int shiftKey) => string.Join("", text.Select(c => Char.IsLetter(c) ? shiftLetter(c, shiftKey) : c));

    public static char shiftLetter(char letter, int shiftKey)
    {
        var letterConst = Char.IsUpper(letter) ? 65 : 97;
        return (char)((letter + shiftKey - letterConst) % 26 + letterConst);
    }
}
