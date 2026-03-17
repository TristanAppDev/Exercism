(ns rna-transcription)

(defn- transcribe
  [nucleotide]
  (case nucleotide
    \G "C"
    \C "G"
    \T "A"
    \A "U"
    (throw (AssertionError. "Illegal nucleotide!"))))

(defn to-rna
  [dna]
  (->> dna
       (map transcribe)
       (apply str)))