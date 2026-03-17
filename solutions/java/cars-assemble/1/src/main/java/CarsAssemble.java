public class CarsAssemble {

    private final int PRODUCTION_PER_HOUR = 221;

    private double normalProductionRate(int speed) {
        return speed * PRODUCTION_PER_HOUR;
    }

    public double productionRatePerHour(int speed) {
        if (speed >= 1 && speed <= 4) {
            return normalProductionRate(speed);
        } else if (speed >= 5 && speed <= 8) {
            return normalProductionRate(speed) * 0.9;
        } else if (speed == 9) {
            return normalProductionRate(speed) * 0.8;
        } else if (speed == 10) {
            return normalProductionRate(speed) * 0.77;
        }
        return 0;
    }

    public int workingItemsPerMinute(int speed) {
        return (int) (productionRatePerHour(speed) / 60);
    }
}
