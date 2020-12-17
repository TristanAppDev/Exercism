pub fn raindrops(n: u32) -> String {
    
    if n % 3 != 0 && n % 5 != 0 && n % 7 != 0 {
        return n.to_string();
    }

    let result = [
        if n % 3 == 0 {"Pling"} else {""},
        if n % 5 == 0 {"Plang"} else {""},
        if n % 7 == 0 {"Plong"} else {""},
    ]
    .join("");

    result
}
