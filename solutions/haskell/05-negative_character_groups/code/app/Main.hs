module Main where

import Data.Char
import Data.List
import Data.Maybe
import System.Environment
import System.Exit

getAllowedCharacters :: String -> String
getAllowedCharacters s = take (length s - 2) $ tail s

getNotAllowedCharacters :: String -> String
getNotAllowedCharacters s = take (length s - 3) $ tail $ tail s

matchPattern :: String -> String -> Bool
matchPattern pattern input = do
  if pattern == "\\d"
    then any isDigit input
    else
      if pattern == "\\w"
        then any (\c -> isAlphaNum c || c == '_') input
        else
          if take 2 pattern == "[^"
            then not $ or $ fmap (`elem` input) (getNotAllowedCharacters pattern)
            else
              if head pattern == '['
                then or $ fmap (`elem` input) (getAllowedCharacters pattern)
                else
                  if length pattern == 1
                    then head pattern `elem` input
                    else error $ "Unhandled pattern: " ++ pattern

main :: IO ()
main = do
  args <- getArgs
  let pattern = args !! 1
  input_line <- getLine

  if head args /= "-E"
    then do
      putStrLn "Expected first argument to be '-E'"
      exitFailure
    else
      if matchPattern pattern input_line
        then exitSuccess
        else exitFailure
