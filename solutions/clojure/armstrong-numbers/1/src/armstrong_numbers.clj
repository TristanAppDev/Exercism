(ns armstrong-numbers)

(defn armstrong? [number]
  (let [num (str number)
        digit-count (count num)]
    (->> num
         (map int)
         (map #(- % 48))
         (map #(apply * (repeat digit-count %)))
         (apply +)
         (= number))))
