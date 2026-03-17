class Bob {

    String hey(String input) {

        input = input.trim();

        if (input.isEmpty()) {
            return "Fine. Be that way!";
        }

        if (isQuestion(input) && isShouting(input)) {
            return "Calm down, I know what I'm doing!";
        }

        if (isQuestion(input)) {
            return "Sure.";
        }

        if (isShouting(input)) {
            return "Whoa, chill out!";
        }

        return "Whatever.";
    }

    private boolean isQuestion(String input) {
        return input.endsWith("?");
    }

    private boolean isShouting(String input) {
        return input.toUpperCase().equals(input) && input.chars().anyMatch(Character::isLetter);
    }

}