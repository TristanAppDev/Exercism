using System;

static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        var idStr = id.HasValue ? $"[{id}] " : "";
        department ??= "OWNER";
        return $"{idStr}{name} - {department.ToUpper()}";
    }
}
