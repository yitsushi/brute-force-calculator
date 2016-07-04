package main

import "flag"

func parseArgs() {
	flag.BoolVar(&cliMode, "c", false, "Enable CLI mode")
	flag.Var(&hashPerSeconds, "r", "rate aka. Hash/sec (available suffix: G M K)\n     \tor use 'auto' to calculate on the fly")

	flag.Parse()
}

func main() {
	parseArgs()

	if cliMode {
		CLIDisplay()
	} else {
		GUIDisplay()
	}
}
