class IsogramChecker {

    boolean isIsogram(String phrase) {
        phrase = phrase.toLowerCase();

        return phrase.chars().filter(Character::isLetter).distinct().count() == phrase.chars()
                .filter(Character::isLetter).count();
    }

}
