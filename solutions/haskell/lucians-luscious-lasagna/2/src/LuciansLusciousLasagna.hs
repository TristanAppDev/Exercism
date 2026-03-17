module LuciansLusciousLasagna (elapsedTimeInMinutes, expectedMinutesInOven, preparationTimeInMinutes) where

-- TODO: define the expectedMinutesInOven constant
expectedMinutesInOven :: Int
expectedMinutesInOven = 40


-- TODO: define the preparationTimeInMinutes function
preparationTimeInMinutes :: Int -> Int
preparationTimeInMinutes layers = layers * 2


-- TODO: define the elapsedTimeInMinutes function
elapsedTimeInMinutes :: Int -> Int -> Int
elapsedTimeInMinutes layers timeSpentInOven = preparationTimeInMinutes layers + timeSpentInOven 