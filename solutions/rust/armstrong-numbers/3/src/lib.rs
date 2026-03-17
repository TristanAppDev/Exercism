pub fn is_armstrong_number(num: u32) -> bool {
    if num < 10 {
        return true;
    }

    let num = num as u64;
    let exp = num.ilog10() + 1;
    let mut result: u64 = 0;
    let mut n = num;

    while n > 0 {
        result += (n % 10).pow(exp);
        n /= 10;
    }

    num == result
}
