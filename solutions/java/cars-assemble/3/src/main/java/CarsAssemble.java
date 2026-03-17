public class CarsAssemble {

    private final int PRODUCTION_PER_HOUR = 221;

    public double productionRatePerHour(int speed) {
        return speed * PRODUCTION_PER_HOUR * successRate(speed);
    }

    private double successRate(int speed) {
        if(speed == 10) return 0.77;
        if(speed == 9) return 0.8;
        if(speed >= 5) return 0.9;
        return 1;
    }

    public int workingItemsPerMinute(int speed) {
        return (int) (productionRatePerHour(speed) / 60);
    }
}