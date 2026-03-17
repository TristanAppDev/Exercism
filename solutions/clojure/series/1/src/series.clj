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
  (cond 
    (= length 0) (vector "")
    (or (empty? string) (< (count string) length)) (vector)
    :else (slices-helper string length [])))
