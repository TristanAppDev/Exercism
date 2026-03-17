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
  (let [[first-el second-el & rest-deck] deck]
    (vec (conj rest-deck first-el second-el))))

(defn discard-top-card
  "Returns a vector containing the first card and
   a vector of the remaining cards in the deck."
  [deck]
  (let [[a & b] deck]
    (if (nil? b)
      (conj [a] nil)
      (conj [a] (vec b)))))

(def face-cards
  ["jack" "queen" "king"])

(defn insert-face-cards
  "Returns the deck with face cards between its head and tail."
  [deck]
  (if (empty? deck)
    (vec face-cards)
    (let [[head & tail] deck]
      (if (nil? head)
        (deck)
        (conj tail "king" "queen" "jack" head)))))


(comment)