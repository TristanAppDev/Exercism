module LeapYear (isLeapYear) where

isLeapYear :: Integer -> Bool
isLeapYear year = not (isDivisible 100) && isDivisible 4 || isDivisible 400
  where
    isDivisible x = year `rem` x == 0