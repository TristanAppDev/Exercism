pub fn factors(n: u64) -> Vec<u64> {    
    let mut result = vec![];
    let mut remainder = n;
    let mut counter = 2;

    while remainder > 1 {
        while remainder % counter == 0 {
            result.push(counter);
            remainder /= counter;
        }
        counter += 1;
    }

    result
}
