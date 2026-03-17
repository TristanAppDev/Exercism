(ns sublist 
  (:import (java.util Collections)))

(defn is-sub-list [list1 list2] 
  (>= (Collections/indexOfSubList list1 list2) 0))

(defn classify [list1 list2]
  (cond
    (and (is-sub-list list1 list2) (= (count list1) (count list2))) :equal
    (and (is-sub-list list1 list2) (> (count list1) (count list2))) :superlist
    (and (is-sub-list list2 list1) (< (count list1) (count list2))) :sublist
    :else :unequal))
  
(classify '(1 2 3 4) '(1 2 4))
