package main

import (
	"fmt"
)

type Letter struct {
	value string
	color string
}

type Guess [5]Letter

type Board struct {
	Rows [6]Guess
}

// TODO: add all letters
var letterOptions = [3]string{"a", "b", "c"}
var colorBlack = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"

func newBoard() Board {
	defaultLetter := Letter{value: "*", color: colorBlack}
	guess := Guess{defaultLetter, defaultLetter, defaultLetter, defaultLetter, defaultLetter}
	board := Board{Rows: [6]Guess{guess, guess, guess, guess, guess, guess}}
	return board
}

// print current state of the board
func (b Board) print() {
	for _, guessRow := range b.Rows {
		for i := 0; i <= len(guessRow)-1; i++ {
			// Print each letter value in appropriate color
			fmt.Print(string(guessRow[i].color), guessRow[i].value, " ")
		}
		fmt.Println()
	}
}

// takes in a word guess
// checks to see if any/all letters of the word are a match
// and returns state
func (b Board) checkGuess(guess string) {
	fmt.Println("You guessed", guess)
	fmt.Println("now want to check this against the word", getNewWord())
}
