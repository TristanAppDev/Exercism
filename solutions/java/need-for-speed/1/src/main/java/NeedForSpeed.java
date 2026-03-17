class NeedForSpeed {
    private int distance;
    private int battery;
    private int speed;
    private int batteryDrain;

    public NeedForSpeed(int speed, int batteryDrain) {
        this.battery = 100;
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return battery == 0;
    }

    public int distanceDriven() {
        return distance;
    }

    public void drive() {
        if (battery - batteryDrain < 0) {
            return;
        }

        battery -= batteryDrain;
        distance += speed;
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

class RaceTrack {
    private final int distance;

    public RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean tryFinishTrack(NeedForSpeed car) {
        
        while (!car.batteryDrained()) {
            car.drive();
        }

        return car.distanceDriven() >= distance;
    }
}
