import java.util.List;
import java.util.stream.Collectors;

class Anagram {

    private final String word;

    public Anagram(String word) {
        this.word = word;
    }

    public List<String> match(List<String> candidates) {
        return candidates
                .stream()
                .filter(other -> !other.equalsIgnoreCase(word))
                .filter(other -> isAnagram(other))
                .collect(Collectors.toList());
    }

    private boolean isAnagram(String other) {
        return word.toLowerCase().chars().sorted().boxed().collect(Collectors.toList())
                .equals(other.toLowerCase().chars().sorted().boxed().collect(Collectors.toList()));
    }

}