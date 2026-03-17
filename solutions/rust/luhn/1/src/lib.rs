/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    if contains_invalid_char(code) || is_single_digit(code)
    {
        return false;
    }

    code.chars()
        .filter_map(|c| c.to_digit(10))
        .rev()
        .enumerate()
        .map(|(index, digit)| if index % 2 == 0 { digit } else { digit * 2 })
        .map(|digit| if digit > 9 { digit - 9 } else { digit })
        .sum::<u32>()
        % 10
        == 0
}

pub fn is_single_digit(code: &str) -> bool {
    code.chars().filter(|c| c.is_digit(10)).take(2).count() <= 1
}

pub fn contains_invalid_char(code: &str) -> bool {
    code.chars().any(|c| !c.is_digit(10) && c != ' ')
}