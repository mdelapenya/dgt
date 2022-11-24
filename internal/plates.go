package internal

import (
	"strconv"
	"strings"
)

// Alphabet is the alphabet used in car plates
const Alphabet = "BCDFGHJKLMNPQRSTVWXYZ"

var lowerCaseAlphabet = strings.ToLower(Alphabet)

// FromPlate returns the index of the plate, and the three characters
func FromPlate(fromPlate string) (initialIndex int, firstChar int, secondChar int, thirdChar int) {
	fromPlate = strings.ToLower(fromPlate)

	initialIndex = 0
	firstChar = 0
	secondChar = 0
	thirdChar = 0

	// the plate must have 7 characters, being the first 4 numbers and the last 3 letters
	plateNumber, err := strconv.Atoi(string(fromPlate[0:4]))
	if err == nil {
		initialIndex = plateNumber
	}

	firstChar = strings.IndexRune(lowerCaseAlphabet, rune(fromPlate[4]))
	secondChar = strings.IndexRune(lowerCaseAlphabet, rune(fromPlate[5]))
	thirdChar = strings.IndexRune(lowerCaseAlphabet, rune(fromPlate[6]))

	return
}
