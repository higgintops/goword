package main

var wordChoices = [...]string{"power", "clock", "trash"}

// TODO: this needs to "randomly" select a word from wordChoices
func getNewWord() string {
	return wordChoices[0]
}

// TODO: also just store which word is being used for the entire duration of a game
