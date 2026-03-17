module LeapYear (isLeapYear) where

isLeapYear :: Integer -> Bool
isLeapYear year = isDivisable 4 && (not (isDivisable 100) || isDivisable 400)
    where
        isDivisable x = year `rem` x == 0