pub fn nth(n: u32) -> u32 {
    (2..)
        .filter(|x| is_prime(x.to_owned()))
        .nth(n as usize)
        .unwrap()
}

fn is_prime(num: u32) -> bool {
    match num {
        0 | 1 => false,
        2 => true,
        _ => (2..=last_divisor(num)).all(|divisor| num % divisor != 0),
    }
}

fn last_divisor(num: u32) -> u32 {
    (num as f64).sqrt().ceil() as u32
}
