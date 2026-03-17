using System;
class RemoteControlCar
{
    private int speed;
    private int batteryDrain;
    private int distance;
    private int battery = 100; 
    public RemoteControlCar(int speed, int BatteryDrain)
    {
        this.speed = speed;
        this.batteryDrain = BatteryDrain;
    }
    public bool BatteryDrained()
    {
        return battery < batteryDrain;
    }
    public int DistanceDriven()
    {
        return distance; 
    }
    public void Drive()
    {
        if (!BatteryDrained())
        {
            this.distance += speed;
            this.battery -= this.batteryDrain;
            Console.WriteLine(this.battery);
        }
    }
    public static RemoteControlCar Nitro()
    {
        return new RemoteControlCar(50, 4);
    }
}

class RaceTrack
{
    private int distance = 0;
    public RaceTrack(int distance)
    {
        this.distance = distance;
    }
    public bool TryFinishTrack(RemoteControlCar car)
    {
        
        while (!car.BatteryDrained() && car.DistanceDriven() < distance)
        {
            car.Drive();
        }
        return car.DistanceDriven() >= distance;
    }
}