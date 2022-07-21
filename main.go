package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const guessCount int = 6

func main() {

	won := false
	word := getNewWord()
	b := newBoard()

	// Each iteration do:
	// 1.) print board
	// 2.) print letter options
	// 3.) ask for input (a word guess)

	for turn := 0; turn <= guessCount; turn++ {
		// clear terminal
		fmt.Print("\033[H\033[2J")

		// Print the game board
		b.print()

		// check win/lose
		if won {
			fmt.Println("~~~~~ YOU WIN!! ~~~~~")
			os.Exit(1)
		} else if turn == 6 {
			fmt.Println("~~~~~ YOU LOST!! ~~~~~ The word was", word)
			os.Exit(1)
		}

		// TODO: print the letter options

		// Ask for user input until it is exactly 5 characters
		guess := ""
		for ok := true; ok; ok = (len(guess) != 5) {
			guess = getUserGuess()
		}

		b = b.checkGuess(guess, turn)
		won = guess == word
	}

}

func getUserGuess() string {
	fmt.Printf("Enter a 5 letter word: ")
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return strings.TrimSuffix(guess, "\n")
}
