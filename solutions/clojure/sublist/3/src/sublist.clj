(ns sublist
  (:require [clojure.set :refer [subset? superset?]]))

(defn classify [list1 list2]
  (cond
    (= list1 list2) :equal
    (subset? list1 list2) :sublist
    (superset? list1 list2) :superlist
    :else :unequal))
