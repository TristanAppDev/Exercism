(ns series
  (:require [clojure.string :as str]))

(defn- slices-helper
  [string length result]
  (if (< (count string) length)
    (vec result)
    (slices-helper
     (str/join "" (rest string))
     length
     (conj result (str/join "" (take length string))))))

(defn slices
  [string length]
  (if (zero? length)
    [""]
    (map #(apply str %) (partition length 1 string))))
