module Pangram (isPangram) where
import Data.Char as Char ( isAsciiLower )
import Data.Function ( (&) )
import Data.Set as Set ( fromList, size )
import Data.Text as Text ( filter, pack, toLower, unpack )

isPangram :: String -> Bool
isPangram text =
  let alphabetLen = 26
   in text
        & Text.pack
        & Text.toLower
        & Text.filter Char.isAsciiLower
        & Text.unpack
        & Set.fromList
        & Set.size
        & (==) alphabetLen