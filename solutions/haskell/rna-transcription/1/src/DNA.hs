module DNA (toRNA) where

toRNA :: String -> Either Char String
toRNA = mapM translateDnaToRna

translateDnaToRna :: Char -> Either Char Char
translateDnaToRna 'G' = Right 'C'
translateDnaToRna 'C' = Right 'G'
translateDnaToRna 'T' = Right 'A'
translateDnaToRna 'A' = Right 'U'
translateDnaToRna  x  = Left x