package main

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
	cliMode        bool       = false
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
