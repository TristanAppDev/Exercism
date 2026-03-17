pub fn reply(message: &str) -> &str {

    let message = message.trim();

    if message.is_empty() {
        return "Fine. Be that way!"
    }

    if is_forceful_question(message) {
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

fn is_forceful_question(message: &str) -> bool {
    is_question(message) && is_forceful(message)
}

fn is_forceful(message: &str) -> bool {
    message.chars().filter(|c| c.is_ascii_alphabetic()).all(|c| c.is_ascii_uppercase()) && has_alphabeticals(message)
}

fn is_question(message: &str) -> bool {
    message.chars().last().unwrap() == '?'
}

fn has_alphabeticals(message: &str) -> bool  {
    message.chars().filter(|c| c.is_ascii_alphabetic()).count() > 0
}