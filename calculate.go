package main

import (
	"fmt"
	"math"
	"strings"
)

func getCharsetForPassword(password string) string {
	var actualCharsetIndex int

	for _, char := range password {
		for charsetIndex, charset := range characterSets {
			if charsetIndex < actualCharsetIndex {
				continue
			}
			if strings.ContainsRune(charset, char) {
				actualCharsetIndex = charsetIndex
				break
			}
		}
	}

	if actualCharsetIndex < 0 {
		actualCharsetIndex = len(characterSets) - 1
	}

	return characterSets[actualCharsetIndex]
}

func calculate(password string) float64 {
	var charset string

	charset = getCharsetForPassword(password)
	fmt.Printf("'%s' with '%s'\n", password, charset)
	var attempts float64 = 0

	passwordLength := len(password)
	charsetLength := len(charset)

	for characterIndexInPassword, character := range password {
		index := strings.IndexRune(charset, character) + 1
		powerOf := passwordLength - characterIndexInPassword - 1

		if powerOf == 0 {
			attempts += float64(index)
		} else {
			attempts += math.Pow(float64(charsetLength), float64(powerOf)) * float64(index)
		}
	}

	return attempts / float64(hashPerSeconds)
}
