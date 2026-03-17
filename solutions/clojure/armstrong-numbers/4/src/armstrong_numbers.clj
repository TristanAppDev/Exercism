(ns armstrong-numbers)

(defn- power
  [base exp]
  (apply * (repeat exp base)))

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