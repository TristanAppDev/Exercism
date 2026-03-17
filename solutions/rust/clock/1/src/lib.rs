use std::fmt;

const MINUTES_PER_HOUR: i32 = 60;
const HOURS_PER_DAY: i32 = 24;
const MINUTES_PER_DAY: i32 = MINUTES_PER_HOUR * HOURS_PER_DAY;

#[derive(Debug, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut minutes = (hours * MINUTES_PER_HOUR + minutes) % MINUTES_PER_DAY;
        minutes = if minutes < 0 {
            minutes + MINUTES_PER_DAY
        } else {
            minutes
        };
        Self { minutes }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(0, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let hours = (self.minutes / MINUTES_PER_HOUR) % HOURS_PER_DAY;
        let minutes = self.minutes % MINUTES_PER_HOUR;
        write!(f, "{:>02}:{:>02}", hours, minutes)
    }
}
