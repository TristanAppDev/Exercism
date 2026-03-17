import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

class IsogramChecker {

    boolean isIsogram(String phrase) {
        return phrase.chars()
                .filter(Character::isLetter)
                .filter(c -> isDuplicate(phrase, c))
                .count() == 0;
    }

    private boolean isDuplicate(String phrase, int c) {
        return Collections.frequency(getListFromString(phrase), c) > 1;
    }

    private List<Integer> getListFromString(String value) {
        return value.toLowerCase().chars().boxed().collect(Collectors.toList());
    }
}