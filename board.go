package main

import (
	"fmt"
)

const guessCount int = 6
const wordLength int = 5

// Color outputs
const colorBlack string = "\033[0m"
const colorRed string = "\033[31m"
const colorGreen string = "\033[32m"

type Letter struct {
	value string
	color string
}

type Guess [wordLength]Letter

type Board struct {
	Rows [guessCount]Guess
}

func newBoard() Board {
	defaultLetter := Letter{value: "*", color: colorBlack}
	guess := Guess{defaultLetter, defaultLetter, defaultLetter, defaultLetter, defaultLetter}
	board := Board{Rows: [guessCount]Guess{guess, guess, guess, guess, guess, guess}}
	return board
}

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
// and returns state of board
func (b Board) checkGuess(userInput string, turn int) Board {
	word := getCurrentWord()
	guess := Guess{}

	for i, l := range userInput {
		color := colorBlack
		// check to see if letter in right placement
		if string(word[i]) == string(l) {
			color = colorGreen
		} else if contains(word, string(l)) {
			color = colorRed
		}

		letter := Letter{value: string(l), color: color}
		guess[i] = letter
	}

	// update board with the guess
	b.Rows[turn] = guess

	return b
}

// TODO:
// need to check Letter of Guess against the word (to check colors)
func contains(word string, l string) bool {
	for i := 0; i <= len(word)-1; i++ {
		if string(word[i]) == l {
			return true
		}
	}

	return false
}
