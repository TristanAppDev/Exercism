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
        & getSubset
        & length
        & (==) textLength

getSubset :: String -> String
getSubset text = [c | c <- alphabet, c `elem` text]