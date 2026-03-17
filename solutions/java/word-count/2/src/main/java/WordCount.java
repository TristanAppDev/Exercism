import java.util.HashMap;
import java.util.Map;

class WordCount {
    public Map<String, Integer> phrase(String input) {

        Map<String, Integer> result = new HashMap<>();
        input = input.toLowerCase().replaceAll("[^a-z1-9']|\\B'|'\\B", " ").trim();
        var words = input.split("\\s+");

        for (String word : words) {
            result.put(word, result.get(word) == null ? 1 : result.get(word) + 1);
        }

        return result;
    }
}
