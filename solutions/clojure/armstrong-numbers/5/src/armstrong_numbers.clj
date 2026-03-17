(ns armstrong-numbers)

(defn- power
  [base exp]
  (apply * (repeat exp base)))

; for practice: tail recursion optimization with 'recur'
(defn- get-digits-rec
  [number]
  (loop [num number
         result '()]
    (if (zero? num)
      result
      (recur (quot num 10) (conj result (rem num 10))))))

; solution
(defn- get-digits
  [number]
  (->> number
       (iterate #(quot % 10))
       (take-while pos?)
       (map #(rem % 10))))

(defn armstrong?
  [number]
  (let [digits (get-digits number)
        digit-count (count digits)]
    (->> digits
         (map #(power % digit-count))
         (apply +)
         (= number))))