# goword
Wordle in Golang

## Rules/Logic (cli version)
* There are 5 blank character spaces ( * * * * * )
* There are 6 chances to guess the word of the day
*  If word guessed correctly: message will display to indicate correctness
*  If word is partially correct:
    * X will indicate a correct letter in a correct placement
    * ? will indicate a correct letter in an incorrect placement
* Each iteration (guess round) a cli version of the guesses will appear as well as which letters are still available to be chosen
    * If a letter already chosen is chosen again: just output message saying so
    * If word is invalid: just output message saying so

* If all 6 rows done and not correct, reveal the word!

