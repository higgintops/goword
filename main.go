package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	won := false
	word := getNewWord()
	b := newBoard()

	for turn := 0; turn <= guessCount; turn++ {
		// clear terminal
		fmt.Print("\033[H\033[2J")

		// Print the game board
		b.print()

		// check win/lose state
		if won {
			fmt.Println("~~~~~ YOU WIN!! ~~~~~")
			os.Exit(1)
		} else if turn == guessCount {
			fmt.Println("~~~~~ YOU LOST!! ~~~~~ The word was", word)
			os.Exit(1)
		}

		// Ask for user input, ensuring word is 5 characters
		guess := ""
		for ok := true; ok; ok = (len(guess) != wordLength) {
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
