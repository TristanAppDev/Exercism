(ns complex-numbers)

(defn real
  [[a b]]
  a)

(defn imaginary
  [[a b]]
  b)

(defn abs
  [[a b]]
  (Math/sqrt (+ (Math/pow a 2) (Math/pow b 2))))

(defn conjugate
  [[a b]]
  (vec [a (* b -1)]))

(defn add
  [[a b] [c d]]
  (vec [(+ a c) (+ b d)]))

(defn sub
  [[a b] [c d]]
  (vec [(- a c) (- b d)]))

(defn mul
  [[a b] [c d]]
  (vec [(- (* a c) (* b d)) (+ (* a d) (* b c))]))

(defn div
  [[a b] [c d]]
  (let [conjugate-cd (conjugate [c d])
        div (first (mul [c d] conjugate-cd))
        numbers (mul [a b] conjugate-cd)]
    (vec [(double (/ (first numbers) div))
          (double (/ (second numbers) div))])))
