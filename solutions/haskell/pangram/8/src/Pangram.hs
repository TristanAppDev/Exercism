module Pangram (isPangram) where

import Data.Char as Char (toLower)

alphabet :: [Char]
alphabet = ['a' .. 'z']

isPangram :: String -> Bool
isPangram text = isInAlphabet (map toLower text)

isInAlphabet :: String -> Bool
isInAlphabet text = all (`elem` text) alphabet