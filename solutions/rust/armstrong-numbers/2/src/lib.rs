pub fn is_armstrong_number(num: u32) -> bool {
    let mut numbers = Vec::new();
    let mut num_temp = num as u64;
    while num_temp > 0 {
        numbers.push(num_temp % 10);
        num_temp = num_temp / 10;
    }

    let mut result: u64 = 0;
    let exp = numbers.len() as u32;

    for n in numbers.to_owned() {
        result += n.pow(exp);

        if result > u32::MAX as u64 {
            return false;
        }
    }

    result as u32 == num
}
