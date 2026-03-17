(ns armstrong-numbers)

(defn- power
  [base exp]
  (apply * (repeat exp base)))

(defn- get-digits
  [num]
  (->> num
       (iterate #(quot % 10))
       (take-while pos?)
       (map #(rem % 10))))

(defn armstrong?
  [number]
  (let [digits (get-digits number)]
    (->> digits
         (map #(power % (count digits)))
         (apply +)
         (= number))))