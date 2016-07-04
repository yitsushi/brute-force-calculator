package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func CLIDisplay() {
	var password string

	password = flag.Arg(0)

	if len(password) < 1 {
		fmt.Print("Password to check: ")
		line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		password = string(line)
	}

	fmt.Printf("Hashes per Seconds: %s\n", hashPerSeconds.String())

	seconds := calculate(password)

	if seconds > OneYearInSeconds {
		years := seconds / OneYearInSeconds
		suffix := ""
		if years > 1000000 {
			years = years / 1000000
			suffix = " million"
			if years > 1000 {
				years = years / 1000
				suffix = " billion"
			}
		}
		fmt.Printf("Years: %.2f%s\n", years, suffix)
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
