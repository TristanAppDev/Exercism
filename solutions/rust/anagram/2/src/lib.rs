use itertools::Itertools;
use std::collections::HashSet;
use unicode_segmentation::UnicodeSegmentation;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let word_lower = word.to_lowercase();
    let word_sorted = sort_text(&word_lower);
    possible_anagrams
        .iter()
        .filter(|pa| is_anagram(&word_sorted, &word_lower, &pa))
        .cloned()
        .collect()
}

fn is_anagram(word_sorted: &str, word_lower: &str, pa: &str) -> bool {
    if word_sorted.len() != pa.len() {
        return false;
    };

    let pa_lower = pa.to_lowercase();

    word_lower != pa_lower && word_sorted == sort_text(&pa_lower)
}

fn sort_text(text: &str) -> String {
    text.graphemes(true).sorted().join("")
}
