pub fn nth(n: u32) -> u32 {
    (2..).filter(|&x| is_prime(x)).nth(n as usize).unwrap()
}

fn is_prime(num: u32) -> bool {
    match num {
        0 | 1 => false,
        2 => true,
        _ => (2..)
            .take_while(|x| x <= &last_divisor(num))
            .all(|divisor| num.rem_euclid(divisor) != 0),
    }
}

fn last_divisor(num: u32) -> u32 {
    (num as f64).sqrt().ceil() as u32
}
