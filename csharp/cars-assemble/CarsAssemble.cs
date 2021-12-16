using System;

static class AssemblyLine
{
    private const int PRODUCTION_PER_HOUR = 221;
    
    private static double NormalProductionRate(int speed, double successRate) => 
        speed * PRODUCTION_PER_HOUR * successRate;

    public static double ProductionRatePerHour_(int speed) {
        if (speed >= 1 && speed <= 4) {
            return NormalProductionRate(speed, 1.0);
        } else if (speed >= 5 && speed <= 8) {
            return NormalProductionRate(speed, 0.9);
        } else if (speed == 9) {
            return NormalProductionRate(speed, 0.8);
        } else if (speed == 10) {
            return NormalProductionRate(speed, 0.77);
        } else {
            return 0;
        }
    }

    public static double ProductionRatePerHour(int speed) => speed switch {
        _ when speed == 9 => NormalProductionRate(speed, 0.8),
        _ when speed == 10 => NormalProductionRate(speed, 0.77),
        _ when speed >= 5 => NormalProductionRate(speed, 0.9),
        _ when speed >= 1 => NormalProductionRate(speed, 1.0),
        _ => 0
    };

    public static int WorkingItemsPerMinute(int speed) => 
        (int) (ProductionRatePerHour(speed) / 60);
}
