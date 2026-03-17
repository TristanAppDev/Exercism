{-# LANGUAGE OverloadedStrings #-}
module Bob (responseFor) where

import qualified Data.Char as C
import qualified Data.Text as T

import Data.Text (Text)

responseFor :: Text -> Text
responseFor text
  | isSilence          = "Fine. Be that way!"
  | isShoutingQuestion = "Calm down, I know what I'm doing!"
  | isShouting         = "Whoa, chill out!"
  | isQuestion         = "Sure."
  | otherwise          = "Whatever."
  where
    isShouting         = T.any C.isAlpha text && T.toUpper text == text
    isQuestion         = T.isSuffixOf "?" (T.stripEnd text)
    isShoutingQuestion = isShouting && isQuestion
    isSilence          = T.null (T.strip text)
