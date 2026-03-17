(ns hamming)

(defn distance_thread_last [strand1 strand2]
  (when (= (count strand1) (count strand2))
    (->>
      (map = strand1 strand2)
      (filter false?)
      (count))))

(defn distance [strand1 strand2]
  (when (= (count strand1) (count strand2))
  (count (filter false? (map = strand1 strand2)))))