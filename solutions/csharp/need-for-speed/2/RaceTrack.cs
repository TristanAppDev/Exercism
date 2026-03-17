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