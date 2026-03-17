(ns nth-prime)

(defn- prime?
  [n]
  (if (= n 2)
    true
    (let [max-n #(inc (Math/sqrt n))]
      (empty? (filter #(zero? (mod n %)) (range 2 (max-n)))))))

(defn nth-prime
  [n]
  (if (> n 0)
    (last (take n (filter prime? (iterate inc 2))))
    (throw (IllegalArgumentException. "no 0th prime"))))