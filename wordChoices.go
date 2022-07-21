package main

import (
	"math/rand"
	"time"
)

// TODO: could read in words from a file
var wordChoices = [...]string{"power", "clock", "trash", "bacon", "cable", "extra", "float", "hello", "jaunt", "knock", "limbo", "quirk", "robot", "shine", "tempo", "twist", "usher", "wager"}
var currentWord string

func getNewWord() string {
	rand.Seed(time.Now().UnixNano())
	max := len(wordChoices)
	randomIndex := rand.Intn(max + 1)
	currentWord = wordChoices[randomIndex]
	return currentWord
}

func getCurrentWord() string {
	return currentWord
}
