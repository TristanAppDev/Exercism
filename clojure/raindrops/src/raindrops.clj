(ns raindrops)

(defn convert [number]
  (cond-> nil
    (zero? (rem number 3)) (str "Pling")
    (zero? (rem number 5)) (str "Plang")
    (zero? (rem number 7)) (str "Plong")
    :default (or (str number))))