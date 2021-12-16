using System;

class RemoteControlCar
{
    private int METERS_PER_USE = 20;
    private int BATTERY_DRAIN_PER_USE = 1;
    private int batteryInPercent = 100;
    private int distanceDrivenInMeters = 0;

    public static RemoteControlCar Buy() => 
        new RemoteControlCar();

    public string DistanceDisplay() => 
        $"Driven {distanceDrivenInMeters} meters";

    public string BatteryDisplay() => 
        (batteryInPercent > 0) ? 
        $"Battery at {batteryInPercent}%" :
        "Battery empty";

    public void Drive()
    {
        if (batteryInPercent == 0) return; 
        batteryInPercent -= BATTERY_DRAIN_PER_USE;
        distanceDrivenInMeters += METERS_PER_USE;
    }
}
