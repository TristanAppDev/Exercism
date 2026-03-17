module PerfectNumbers (classify, Classification (..)) where

data Classification = Deficient | Perfect | Abundant deriving (Eq, Show)

classify :: Int -> Maybe Classification
classify num
  | num < 1 = Nothing
  | x < num = Just Deficient
  | x > num = Just Abundant
  | otherwise = Just Perfect
  where
    x = sumOfFactors num

sumOfFactors :: Int -> Int
sumOfFactors num = sum (factors num)

factors :: Int -> [Int]
factors number =
  [divisor | divisor <- [1 .. (number `quot` 2)], number `rem` divisor == 0]