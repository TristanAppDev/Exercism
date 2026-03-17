module Hamming (distance) where

distance :: String -> String -> Maybe Int
distance xs ys =
  if equalLength xs ys
    then Just (distance' xs ys)
    else Nothing

distance' :: String -> String -> Int
distance' xs ys
    | null xs = 0
    | head xs == head ys = distance' (tail xs) (tail ys)
    | otherwise = 1 + distance' (tail xs) (tail ys)

equalLength :: String -> String -> Bool
equalLength xs ys = length xs == length ys