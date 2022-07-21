package main

var wordChoices = [...]string{"power", "clock", "trash"}
var currentWord string

// TODO: this needs to "randomly" select a word from wordChoices
func getNewWord() string {
	currentWord = wordChoices[0]
	return currentWord
}

func getCurrentWord() string {
	return currentWord
}
