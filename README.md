# goword
Wordle in Golang

## Rules/Logic (cli version)
* There are 5 blank character spaces ( * * * * * )
* There are 6 chances to guess the word of the day
*  If word guessed correctly: message will display to indicate correctness
*  If word is partially correct:
    * A green letter indicates a letter in the correct placement
    * A red letter indicates a letter an incorrect placement, but still on the board
* Each iteration of guessing:
    * No validation on previously guessed letters/words
    * No validation on invalid words/spelling
    * Only validation is input must be 5 characters in length

* If all 6 rows done and not correct, reveal the word!

