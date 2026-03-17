module Bob (responseFor) where

import qualified Data.Char as C
import qualified Data.Text as T

import Data.Text (Text)

responseFor :: Text -> Text
responseFor text
  | isSilence          = T.pack "Fine. Be that way!"
  | isShoutingQuestion = T.pack "Calm down, I know what I'm doing!"
  | isShouting         = T.pack "Whoa, chill out!"
  | isQuestion         = T.pack "Sure."
  | otherwise          = T.pack "Whatever."
  where
    isShouting         = T.any C.isAlpha text && T.toUpper text == text
    isQuestion         = T.isSuffixOf (T.pack "?") (T.stripEnd text)
    isShoutingQuestion = isShouting && isQuestion
    isSilence          = T.null (T.strip text)
