(ns perfect-numbers)

(defn factors-of
  [num]
  (filter #(zero? (rem num %)) (range 1 (inc (/ num 2)))))

(defn classify
  [num]
  (let [as (reduce + (factors-of num))]
    (cond
      (neg? num) (throw (IllegalArgumentException. "Number shouldn't be negative."))
      (< as num) :deficient
      (> as num) :abundant
      (= as num) :perfect)))

