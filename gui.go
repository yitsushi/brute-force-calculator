package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
)

var (
	password string = "password"

	passwordInput             *ui.Entry
	passwordCrackTimeLabel    *ui.Label
	passwordCrackTimeTemplate string = "%s: %.2f%s"
)

func GUIDisplay() {
	ui.Main(generateMainWindow)
}

func generateMainWindow() {
	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	box.Append(generateHashPowerSide(), true)
	box.Append(generateCalculatorSide(), true)

	window := ui.NewWindow("Brute-Force Calculator", 500, 200, false)
	window.SetChild(box)
	window.SetMargined(true)

	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	window.Show()
}

// Calculator
func generateHashPowerSide() ui.Control {
	hashPowerLabelTemplate := "Current HashPower: %sH/s"
	currentHashPowerLabel := ui.NewLabel(
		fmt.Sprintf(hashPowerLabelTemplate, hashPerSeconds.String()),
	)
	calculatePowerButton := ui.NewButton("Calculate my HashPower")

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(currentHashPowerLabel, false)
	box.Append(calculatePowerButton, false)

	calculatePowerButton.OnClicked(func(*ui.Button) {
		valueChannel := make(chan HashPerSec)
		progressChannel := make(chan int)

		calculatePowerButton.Disable()

		time.Sleep(100 * time.Millisecond)

		go func() {
			for {
				select {
				case progressValue := <-progressChannel:
					{
						updateLabelText(
							currentHashPowerLabel,
							fmt.Sprintf("Calculating HashPower: %d%%", progressValue),
						)
						time.Sleep(50 * time.Millisecond)
					}
				case finalValue := <-valueChannel:
					{
						hashPerSeconds = finalValue

						currentHashPowerLabel.SetText(
							fmt.Sprintf(hashPowerLabelTemplate, hashPerSeconds.String()),
						)
						calculatePowerButton.Enable()

						updateCrackingTime()
						return
					}
				default:
					{
						time.Sleep(50 * time.Millisecond)
					}
				}
			}
		}()

		go messureMD5HashPerSec(progressChannel, valueChannel)
	})

	return box
}

// Password check
func generateCalculatorSide() ui.Control {
	passwordInput = ui.NewEntry()
	passwordInput.SetText(password)
	passwordCrackTimeLabel = ui.NewLabel(
		fmt.Sprintf(passwordCrackTimeTemplate, "Seconds", 0.0, ""),
	)

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(ui.NewLabel("Password to check:"), false)
	box.Append(passwordCrackTimeLabel, false)
	box.Append(passwordInput, false)

	passwordInput.OnChanged(func(item *ui.Entry) {
		password = item.Text()

		updateCrackingTime()
	})

	updateCrackingTime()

	return box
}

func updateCrackingTime() {
	period, value, suffix := calculateTimeForGUI()

	updateLabelText(
		passwordCrackTimeLabel,
		fmt.Sprintf(passwordCrackTimeTemplate, period, value, suffix),
	)
}

func updateLabelText(label *ui.Label, text string) {
	label.SetText(text)
}

func calculateTimeForGUI() (period string, value float64, suffix string) {
	seconds := calculate(password)
	value = seconds
	period = "Seconds"
	suffix = ""

	if seconds > OneYearInSeconds {
		period = "Years"
		value = seconds / OneYearInSeconds

		if value > 1000000 {
			value = value / 1000000
			suffix = " million"

			if value > 1000 {
				value = value / 1000
				suffix = " billion"
			}
		}
	} else if seconds > OneDayInSeconds {
		period = "Days"
		value = seconds / OneDayInSeconds
	} else if seconds > OneHourInSeconds {
		period = "Hours"
		value = seconds / OneHourInSeconds
	} else if seconds > OneMinuteInSeconds {
		period = "Minutes"
		value = seconds / OneMinuteInSeconds
	}

	return
}
