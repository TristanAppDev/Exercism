(ns series)

(defn slices
  [string length]
  (if (zero? length)
    [""]
    (->> string
         (partition length 1)
         (mapv (partial apply str)))))
