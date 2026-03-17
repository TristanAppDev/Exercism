// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

const CARS_PER_HOUR: f64 = 221.0;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let cars_produced: f64 = CARS_PER_HOUR * speed as f64;

    match speed {
        9 | 10 => cars_produced * 0.77,
        5.. => cars_produced * 0.9,
        1.. => cars_produced,
        _ => 0.0,
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}
