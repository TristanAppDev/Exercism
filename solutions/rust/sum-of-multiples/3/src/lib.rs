pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    (1..limit)
        .into_iter()
        .filter(|&num| {
            factors
                .iter()
                .any(|&factor| is_positive(factor) && is_evenly_divisible(num, factor))
        })
        .sum::<u32>()
}

pub fn is_positive(factor: u32) -> bool {
    factor > 0
}

pub fn is_evenly_divisible(num: u32, factor: u32) -> bool {
    num % factor == 0
}
