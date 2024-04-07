package utils

import (
	"math/rand"
)

// Function to generate a random 10-letter word
func GenerateWord(l int) string {
	// Pool of lowercase letters
	letters := "abcdefghijklmnopqrstuvwxyz"

	word := make([]byte, l)
	for i := range word {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}
