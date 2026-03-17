(ns perfect-numbers)

(defn factors-of
  [num]
  (->>
   (range 1 (inc (/ num 2)))
   (filter #(zero? (rem num %)))))

(defn classify
  [num]
  (if (neg? num)
    (throw (IllegalArgumentException. "Number shouldn't be negative."))
    (let [as (apply + (factors-of num))]
      (cond
        (< as num) :deficient
        (> as num) :abundant
        (= as num) :perfect))))
