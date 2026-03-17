using System;

class RemoteControlCar
{
    private int METERS_PER_USE = 20;
    private int BATTERY_DRAIN_PER_USE = 1;
    private int battery_in_percent = 100;
    private int meters_driven = 0;

    public static RemoteControlCar Buy() => 
        new RemoteControlCar();

    public string DistanceDisplay() => 
        $"Driven {meters_driven} meters";

    public string BatteryDisplay() => 
        (battery_in_percent > 0) ? 
        $"Battery at {battery_in_percent}%" :
        "Battery empty";

    public void Drive()
    {
        if (battery_in_percent == 0) return; 
        battery_in_percent -= BATTERY_DRAIN_PER_USE;
        meters_driven += METERS_PER_USE;
    }
}
