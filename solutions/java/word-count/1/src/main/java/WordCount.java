import java.util.HashMap;
import java.util.Map;

class WordCount {
    public Map<String, Integer> phrase(String input) {

        Map<String, Integer> result = new HashMap<>();
        input = input.toLowerCase().replaceAll("[^a-z1-9']|\\B'|'\\B", " ").trim();
        var words = input.split("\\s+");

        for (String word : words) {
            var wordCount = result.get(word);
            result.put(word, wordCount == null ? 1 : wordCount + 1);
        }

        return result;
    }
}
