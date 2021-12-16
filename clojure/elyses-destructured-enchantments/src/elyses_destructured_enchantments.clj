(ns elyses-destructured-enchantments)

(defn first-card
  "Returns the first card from deck."
  [deck]
  (first deck))

(defn second-card
  "Returns the second card from deck."
  [deck]
  (let [[_ second] deck]
    second))

(defn swap-top-two-cards
  "Returns the deck with first two items reversed."
  [deck]
  (let [[first second & rest] deck]
    (vec (conj rest first second))))

(defn discard-top-card
  "Returns a vector containing the first card and
   a vector of the remaining cards in the deck."
  [deck]
  (let [[head & tail] deck]
    [head tail]))

(def face-cards
  ["jack" "queen" "king"])

(defn insert-face-cards
  "Returns the deck with face cards between its head and tail."
  [deck]
  (let [[head & tail] deck]
    (vec (remove nil? (flatten [head face-cards tail])))))

(comment)