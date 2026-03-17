module PerfectNumbers (classify, Classification (..)) where

data Classification = Deficient | Perfect | Abundant deriving (Eq, Show)

classify :: Int -> Maybe Classification
classify num =
    if num < 1
    then Nothing
    else classify' num

classify' :: Int -> Maybe Classification
classify' num
    | factorsSum num < num = Just Deficient
    | factorsSum num > num = Just Abundant
    | otherwise = Just Perfect

factorsSum :: Integral a => a -> a
factorsSum number = sum (factors number)

factors :: Integral a => a -> [a]
factors number =
    [divisor | divisor <- [1 .. (number `quot` 2)], number `rem` divisor == 0]