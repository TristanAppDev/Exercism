(ns leap)

(defn leap-year? [year]
  (if
   (and (== (rem year 4) 0)
        (or (not (= (rem year 100) 0)) (== (rem year 400) 0)))
    true
    false))