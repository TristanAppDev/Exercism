(ns series)

; for practise
(defn slices
  [string length]
  (distinct
   (mapv
    (partial apply str)
    (partition length 1 string))))

; more readable
(defn slices
  [string length]
  (->> string
       (partition length 1)
       (mapv (partial apply str))
       distinct))
