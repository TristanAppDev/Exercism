(ns log-levels
  (:require [clojure.string :as str]))

(defn message
  "Takes a string representing a log line
   and returns its message with whitespace trimmed."
  [text]
  (str/trim (last (str/split text #":"))))

(defn log-level
  "Takes a string representing a log line
   and returns its level in lower-case."
  [text]
  (str/lower-case (subs (first (str/split text #"]")) 1)))

(defn reformat
  "Takes a string representing a log line and formats it
   with the message first and the log level in parentheses."
  [text]
  (str (message text) " (" (log-level text) ")"))
