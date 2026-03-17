(ns nth-prime)

(defn- prime?
  [n]
  (or (= n 2)
      (let [max-n #(inc (Math/sqrt n))]
        (empty? (filter #(zero? (rem n %)) (range 2 (max-n)))))))

(defn nth-prime
  [n]
  (if (> n 0)
    (last (take n (filter prime? (iterate inc 2))))
    (throw (IllegalArgumentException. "no 0th prime"))))