using System;

static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        var idStr = id ?? $"[{id}] ";
        department ??= "OWNER";
        return $"{id}{name} - {department.ToUpper()}";

    }
}
