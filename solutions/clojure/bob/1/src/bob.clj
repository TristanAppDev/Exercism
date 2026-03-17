(ns bob
  (:require [clojure.string :as str]))

(defn- has-letter? [text] (some #(Character/isLetter (char %)) text))
(defn- silence? [text] (str/blank? text))
(defn- question? [text] (= \? (last text)))
(defn- shouting? [text] (and (= text (str/upper-case text)) (has-letter? text)))
(defn- shouting-question? [text] (and (shouting? text) (question? text)))

(defn response-for
  [input]
  (let [text (str/trimr input)]
    (cond
      (shouting-question? text) "Calm down, I know what I'm doing!"
      (silence? text) "Fine. Be that way!"
      (shouting? text) "Whoa, chill out!"
      (question? text) "Sure."
      :else "Whatever.")))