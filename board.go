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
// var letterOptions = [3]string{"a", "b", "c"}
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

	// TODO: return b and guess and check win in main
	return b
}

// Helper function
func contains(word string, l string) bool {
	for i := 0; i <= len(word)-1; i++ {
		if string(word[i]) == l {
			return true
		}
	}

	return false
}

// func (g Guess) toString() string {
// 	s := ""
// 	for _, l := range g {
// 		s += l.value
// 	}
// 	return s
// }
