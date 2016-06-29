package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HashPerSec uint64

func (hsp *HashPerSec) String() string {
	if *hsp < 1000 {
		return fmt.Sprintf("%d", uint64(*hsp))
	} else if *hsp < 1000*1000 {
		return fmt.Sprintf("%dK", uint64(*hsp/(1000)))
	} else if *hsp < 1000*1000*1000 {
		return fmt.Sprintf("%dM", uint64(*hsp/(1000*1000)))
	} else {
		return fmt.Sprintf("%dG", uint64(*hsp/(1000*1000*1000)))
	}
}

func (hsp *HashPerSec) Set(value string) error {
	simpleValue := value[0 : len(value)-1]
	suffix := strings.ToLower(value[len(value)-1:])

	numericValue, err := strconv.ParseUint(simpleValue, 10, 64)
	if err != nil {
		return err
	}

	if suffix == "k" {
		*hsp = HashPerSec(numericValue * uint64(1000))
	} else if suffix == "m" {
		*hsp = HashPerSec(numericValue * uint64(1000*1000))
	} else if suffix == "g" {
		*hsp = HashPerSec(numericValue * uint64(1000*1000*1000))
	} else {
		numericValue, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}

		*hsp = HashPerSec(numericValue)
	}

	return nil
}
