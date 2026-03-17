(ns difference-of-squares)

(defn sum-of-squares [n]
  (quot (* n (inc n) (inc (* 2 n))) 6))

(defn square-of-sum [n]
  (let [n (quot (* n (inc n)) 2)]
    (* n n)))

(defn difference [n]
  (- (square-of-sum n) (sum-of-squares n)))