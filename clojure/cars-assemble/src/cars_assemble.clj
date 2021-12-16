(ns cars-assemble)

(def production-per-hour 221)

(defn- success-rate
  [speed]
  (cond (= speed 10) 0.77
        (= speed 9) 0.8
        (>= speed 5) 0.9
        :else 1.0))

(defn production-rate
  "Returns the assembly line's production rate per hour,
   taking into account its success rate"
  [speed]
  (* (* production-per-hour speed) (success-rate speed)))

(defn working-items
  "Calculates how many working cars are produced per minute"
  [speed]
  (int (quot (production-rate speed) 60)))
