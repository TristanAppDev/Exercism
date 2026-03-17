pub fn is_armstrong_number(num: u32) -> bool {
    let mut numbers = Vec::new();
    let mut num_temp = num;
    while num_temp > 0 {
        numbers.push(num_temp % 10);
        num_temp = num_temp / 10;
    }
    numbers.to_owned().into_iter().map(|n| u32::pow(n, numbers.len() as u32)).sum::<u32>() == num
}
