package main

import (
	"bufio"
	"fmt"
	"os"
)

const guessCount int = 6

func main() {

	word := getNewWord()
	b := newBoard()

	// Each iteration do:
	// 1.) print board
	// 2.) print letter options
	// 3.) ask for input (a word guess)

	fmt.Println("The word:", word)
	for i := 1; i <= guessCount; i++ {
		// clear terminal
		//fmt.Print("\033[H\033[2J")

		// Print the game board
		b.print()

		// print the letter options
		fmt.Println("TODO: print letter options")

		// Ask for user input
		fmt.Printf("Enter a word guess: ")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')

		b.checkGuess(guess)
	}

}
