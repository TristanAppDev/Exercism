(ns bob
  (:require [clojure.string :as str]))

(def silence? str/blank?)

(defn- upper?
  [text]
  (= (str/upper-case text) text))

(defn- question?
  [text]
  (str/ends-with? text "?"))

(defn- has-letter?
  [text]
  (some #(Character/isLetter %) text))

(defn- shouting?
  [text]
  (and (upper? text) (has-letter? text)))

(defn- shouting-question?
  [text]
  (and (shouting? text) (question? text)))

(defn response-for
  [input]
  (let [text (str/trimr input)]
    (cond
      (shouting-question? text) "Calm down, I know what I'm doing!"
      (silence? text) "Fine. Be that way!"
      (shouting? text) "Whoa, chill out!"
      (question? text) "Sure."
      :else "Whatever.")))