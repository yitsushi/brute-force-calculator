package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
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
	if value == "auto" {
		if debugMode {
			fmt.Println("Calculate HasPerSec value based on your computation power with MD5...")
		}
		*hsp = messureMD5HashPerSec()
		return nil
	}

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

func messureMD5HashPerSec() HashPerSec {
	data := []byte("These pretzels are making me thirsty.")
	counter := 0
	max := 100000000

	startTime := time.Now().UnixNano()
	for {
		md5.Sum(data)
		counter++
		if counter >= max {
			break
		}
		if debugMode {
			if counter%1000000 == 0 {
				fmt.Printf(".")
			}
		}
	}
	endTime := time.Now().UnixNano()

	durationInSeconds := int64((endTime - startTime) / int64(time.Second))

	if debugMode {
		fmt.Printf("\n")
		fmt.Printf(
			"%d hashes over %d seconds => %d\n",
			counter,
			durationInSeconds,
			int64(counter)/durationInSeconds,
		)
	}

	return HashPerSec(int64(counter) / durationInSeconds)
}
