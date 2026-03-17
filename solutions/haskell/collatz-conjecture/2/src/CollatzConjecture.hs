module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz num
  | num <= 0 = Nothing
  | otherwise = Just (collatz' num 0)

collatz' :: Integer -> Integer -> Integer
collatz' num step
  | num == 1 = step
  | even num = collatz' (num `div` 2) (step + 1)
  | otherwise = collatz' (3 * num + 1) (step + 1)