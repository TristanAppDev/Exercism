(ns series)

(defn slices
  [string length]
  (->> string
       (partition length 1)
       (mapv (partial apply str))
       distinct))
