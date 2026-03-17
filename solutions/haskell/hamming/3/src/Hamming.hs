module Hamming (distance) where

distance :: String -> String -> Maybe Int
distance xs ys
  | equalLength xs ys = Just (distance' xs ys)
  | otherwise = Nothing

distance' :: String -> String -> Int
distance' xs ys = length $ filter id $ zipWith (/=) xs ys

equalLength :: String -> String -> Bool
equalLength xs ys = length xs == length ys