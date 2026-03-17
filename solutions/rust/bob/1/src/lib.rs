pub fn reply(message: &str) -> &str {
    if message.trim() == "" {
        return "Fine. Be that way!"
    }

    if is_question(message) && is_forceful(message) {
        return "Calm down, I know what I'm doing!"
    }

    if is_question(message) {
        return "Sure."
    }

    if is_forceful(message) {
        return "Whoa, chill out!"
    }

    "Whatever."
}

fn is_forceful(message: &str) -> bool {
    message.trim_end().chars().filter(|c| c.is_ascii_alphabetic()).all(|c| c.is_ascii_uppercase()) && has_alphabeticals(message)
}

fn is_question(message: &str) -> bool {
    message.trim_end().chars().last().unwrap() == '?'
}

fn has_alphabeticals(message: &str) -> bool  {
    message.trim_end().chars().filter(|c| c.is_ascii_alphabetic()).count() > 0
}