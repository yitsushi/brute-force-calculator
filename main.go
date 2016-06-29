package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	// Use 1 billion by default
	// that means 1GH/s
	DefaultHashPerSecond = HashPerSec(1000000000)
	OneYearInSeconds     = 60 * 60 * 24 * 365
	OneDayInSeconds      = 60 * 60 * 24
	OneHourInSeconds     = 60 * 60
	OneMinuteInSeconds   = 60
)

var (
	debugMode      bool       = false
	hashPerSeconds HashPerSec = DefaultHashPerSecond
	characterSets  []string   = []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz0123456789",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]\"{}|;':,./<>?`~ ",
	}
)

func main() {
	var password string

	flag.BoolVar(&debugMode, "d", false, "Enable debug mode")
	flag.Var(&hashPerSeconds, "r", "rate aka. Hash/sec (available suffix: G M K)")

	flag.Parse()

	password = flag.Arg(0)

	if len(password) < 1 {
		line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		password = string(line)
	}

	if debugMode {
		fmt.Printf("Hashes per Seconds: %s\n", hashPerSeconds.String())
	}

	display(calculate(password))
}

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
	if debugMode {
		fmt.Printf("'%s' with '%s'\n", password, charset)
	}

	var attempts uint64 = 0

	passwordLength := len(password)
	charsetLength := len(charset)

	for characterIndexInPassword, character := range password {
		index := strings.IndexRune(charset, character) + 1
		powerOf := passwordLength - characterIndexInPassword - 1

		if powerOf == 0 {
			attempts += uint64(index)
		} else {
			attempts += uint64(math.Pow(float64(charsetLength), float64(powerOf)) * float64(index))
		}
	}

	return float64(attempts / uint64(hashPerSeconds))
}

func display(seconds float64) {
	if seconds > OneYearInSeconds {
		fmt.Printf("Years: %.2f\n", seconds/OneYearInSeconds)
	} else if seconds > OneDayInSeconds {
		fmt.Printf("Days: %.2f\n", seconds/OneDayInSeconds)
	} else if seconds > OneHourInSeconds {
		fmt.Printf("Hours: %.2f\n", seconds/OneHourInSeconds)
	} else if seconds > OneMinuteInSeconds {
		fmt.Printf("Minutes: %.2f\n", seconds/OneMinuteInSeconds)
	} else {
		fmt.Printf("Seconds: %.2f\n", seconds)
	}
}
