package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	showHelp := flag.Bool("h", false, "Show help message.")
	goToBackground := flag.Bool("b", false, "Go to background after startup.")
	quiet := flag.Bool("q", false, "Turn off output.")
	inputFile := flag.String("i", "", "Input file to read URLs from.")
	outputFile := flag.String("o", "", "Log all messages to provided file.")
	outputDirectory := flag.String("d", cwd, "Directory to save downloaded files to. Default value is the current working directory.")
	numTries := flag.Int("t", 1, "Number of retries for each download. Specify 0 for infinite retrying.")
	concurrencyLevel := flag.Int("c", runtime.GOMAXPROCS(0), "Number of concurrent downloads allowed. Default value is GOMAXPROCS (nproc, if unset).")
	flag.Parse()

	if *showHelp {
		fmt.Printf("Usage: %s [OPTION]... [URL]...\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if *goToBackground {
		// Move process to background.
	}

	if *quiet {
		// Redirect stdout to /dev/null.
	}

	if *outputFile != "" {
		// Redirect stdout to output file.
	}

	if *inputFile != "" {
		// Get all download links from file.
	} else {
		// Get all download links from flag.Args().
	}

	fmt.Println(*numTries)
	fmt.Println(*goToBackground)
	fmt.Println(*quiet)
	fmt.Println(*inputFile)
	fmt.Println(*outputFile)
	fmt.Println(*outputDirectory)
	fmt.Println(*concurrencyLevel)
}
