using System.Collections.Generic;

public class SpaceAge
{
    private readonly Dictionary<string, double> PlanetFactors = new Dictionary<string, double>
    {
      {"Mercury", 0.2408467},
      {"Venus", 0.61519726},
      {"Earth", 1.0},
      {"Mars", 1.8808158},
      {"Jupiter", 11.862615},
      {"Saturn", 29.447498},
      {"Uranus", 84.016846},
      {"Neptune", 164.79132}
    };

    private readonly double earthYearInSeconds = 31557600.0;
    private int seconds;

    public double EarthYearInSeconds => earthYearInSeconds;
    public int Seconds { get => seconds; set => seconds = value; }


    public SpaceAge(int seconds)
    {
        this.Seconds = seconds;
    }

    private double calcYears(double planetFactor) => Seconds / EarthYearInSeconds / planetFactor;

    public double OnEarth()
    {
        return calcYears(PlanetFactors["Earth"]);
    }

    public double OnMercury()
    {
        return calcYears(PlanetFactors["Mercury"]);
    }

    public double OnVenus()
    {
        return calcYears(PlanetFactors["Venus"]);
    }

    public double OnMars()
    {
        return calcYears(PlanetFactors["Mars"]);
    }

    public double OnJupiter()
    {
        return calcYears(PlanetFactors["Jupiter"]);
    }

    public double OnSaturn()
    {
        return calcYears(PlanetFactors["Saturn"]);
    }

    public double OnUranus()
    {
        return calcYears(PlanetFactors["Uranus"]);
    }

    public double OnNeptune()
    {
        return calcYears(PlanetFactors["Neptune"]);
    }
}