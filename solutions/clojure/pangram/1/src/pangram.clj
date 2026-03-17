(ns pangram
  (:require [clojure.string :as string]))

(defn pangram? [sentence]
  (->> (string/lower-case sentence)
       (re-seq #"[a-z]")
       (frequencies)
       (count)
       (= 26)))