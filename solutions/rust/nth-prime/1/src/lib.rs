pub fn nth(n: u32) -> u32 {
    let mut prime_counter: u32 = 0;
    let mut current_num: u32 = 1;

    while prime_counter <= n {
        current_num += 1;
        if is_prime(current_num) {
            prime_counter += 1;
        }
    }
    current_num
}

fn is_prime(num: u32) -> bool {
    if num < 2 {
        return false;
    }

    let last_num_to_check = ((num as f64).sqrt() as u32) + 1;

    for i in 2..last_num_to_check {
        if num % i == 0 {
            return false;
        }
    }
    true
}
