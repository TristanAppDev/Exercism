module LeapYear (isLeapYear) where

isLeapYear :: Integer -> Bool
isLeapYear year = isDivisable 400 || isDivisable 4 && not (isDivisable 100)
    where
        isDivisable x = year `mod` x == 0