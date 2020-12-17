module Pangram (isPangram) where

import Data.Char as Char (toLower)

alphabet :: [Char]
alphabet = ['a' .. 'z']

isPangram :: String -> Bool
isPangram text = alphabet `isSubsetOf` map toLower text

isSubsetOf :: String -> String -> Bool
isSubsetOf characters text = all (`elem` text) characters