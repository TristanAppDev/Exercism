#[derive(Debug)]
pub struct Duration(f64);

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration(s as f64)
    }
}

pub trait Planet {
    const PLANET_FACTOR: f64;
    const EARTH_YEARS_IN_SECONDS: f64 = 31557600.0;

    fn years_during(d: &Duration) -> f64 {
        d.0 / (Self::EARTH_YEARS_IN_SECONDS * Self::PLANET_FACTOR)
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

impl Planet for Mercury {
    const PLANET_FACTOR: f64 = 0.2408467;
}
impl Planet for Venus {
    const PLANET_FACTOR: f64 = 0.61519726;
}
impl Planet for Earth {
    const PLANET_FACTOR: f64 = 1.0;
}
impl Planet for Mars {
    const PLANET_FACTOR: f64 = 1.8808158;
}
impl Planet for Jupiter {
    const PLANET_FACTOR: f64 = 11.862615;
}
impl Planet for Saturn {
    const PLANET_FACTOR: f64 = 29.447498;
}
impl Planet for Uranus {
    const PLANET_FACTOR: f64 = 84.016846;
}
impl Planet for Neptune {
    const PLANET_FACTOR: f64 = 164.79132;
}
