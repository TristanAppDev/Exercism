module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz n
  | isNegativeOrZero n = Nothing
  | otherwise = Just (collatz' n)

collatz' :: Integer -> Integer
collatz' n
  | n == 1 = 0
  | even n = 1 + collatz' (evenOperation n)
  | otherwise = 1 + collatz' (oddOperation n)

evenOperation :: Integer -> Integer
evenOperation n = n `div` 2

oddOperation :: Integer -> Integer
oddOperation n = 3 * n + 1

isNegativeOrZero :: Integer -> Bool
isNegativeOrZero n = n <= 0