module Pangram (isPangram) where
import Data.Char as Char ( toLower )
import Data.Function ((&))

alphabet :: [Char]
alphabet = ['a' .. 'z']

isPangram :: String -> Bool
isPangram text = 
    let textLength = length alphabet
      in text
        & map toLower 
        & clean 
        & intersect
        & length
        & (==) textLength

intersect :: String -> String
intersect text = [c | c <- alphabet, c `elem` text]

clean :: String -> String
clean s = [c | c <- s, c `elem` alphabet]