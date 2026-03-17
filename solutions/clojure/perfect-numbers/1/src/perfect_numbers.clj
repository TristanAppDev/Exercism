(ns perfect-numbers)

(defn factors-of
  [n]
  (filter #(zero? (rem n %)) (range 1 (inc (/ n 2)))))

(defn aliquot-sum
  [num]
  (reduce + (factors-of num)))

(defn classify
  [num]
  (cond
    (< num 0) (throw (IllegalArgumentException. "Number shouldn't be negative."))
    (< (aliquot-sum num) num) :deficient
    (> (aliquot-sum num) num) :abundant
    (= (aliquot-sum num) num) :perfect))
