(ns pangram
  (:require [clojure.string :as string]))

(defn pangram? [sentence]
  (->> (string/lower-case sentence)
       (filter Character/isLetter)
       (frequencies)
       (count)
       (= 26)))

(defn pangram-? [sentence]
  (->> (string/lower-case sentence)
       (re-seq #"[a-z]")
       (frequencies)
       (count)
       (= 26)))