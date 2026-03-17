use std::collections::HashSet;
use itertools::Itertools; 
use unicode_segmentation::UnicodeSegmentation;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    possible_anagrams
        .iter()
        .filter(|pa| is_anagram(word, pa) && is_not_equal(word, pa))
        .cloned()
        .collect()
}

pub fn is_not_equal(word: &str, possible_anagram: &str) -> bool {
    !word.to_lowercase().eq(&possible_anagram.to_lowercase())
}

pub fn is_anagram(word: &str, possible_anagram: &str) -> bool {
    word.to_lowercase().graphemes(true).sorted().join("") == possible_anagram.to_lowercase().graphemes(true).sorted().join("")
}
