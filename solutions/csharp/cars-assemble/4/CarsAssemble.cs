using System;

static class AssemblyLine
{
    private const int PRODUCTION_PER_HOUR = 221;

    public static double ProductionRatePerHour(int speed) => 
        speed * PRODUCTION_PER_HOUR * SuccessRate(speed);

    private static double SuccessRate(int speed) => speed switch {
        10 =>  0.77,
        9 => 0.8,
        >= 5 => 0.9,
        _ => 1.0,
    };

    public static int WorkingItemsPerMinute(int speed) => 
        (int) (ProductionRatePerHour(speed) / 60);
}
