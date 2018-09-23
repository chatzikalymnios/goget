package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/chatzikalymnios/goget/internal/utils"
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

	out := os.Stdout
	if *outputFile != "" {
		f, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE,0660)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		out = f
	}

	if *concurrencyLevel < 1 {
		fmt.Fprintf(out, "Error: concurrency level can't be less than 1\n")
		os.Exit(1)
	}

	if *goToBackground {
		// Move process to background.
	}

	if *quiet {
		// Redirect stdout to /dev/null.
	}

	lines := flag.Args()
	if *inputFile != "" {
		lines, err = utils.ReadLines(*inputFile)
		if err != nil {
			fmt.Fprintf(out, "Error reading file: %v\n", err)
			os.Exit(1)
		}
	}

	urls, err := utils.StringToURL(lines)
	if err != nil {
		fmt.Fprintf(out, "Error parsing URLs: %v\n", err)
		os.Exit(1)
	}

	utils.DownloadURLs(urls, *concurrencyLevel, *outputDirectory, out, *quiet)

	//fmt.Println(urls)

	fmt.Println(*numTries)
	fmt.Println(*goToBackground)
	fmt.Println(*quiet)
	//fmt.Println(*inputFile)
	//fmt.Println(*outputFile)
	//fmt.Println(*outputDirectory)
	//fmt.Println(*concurrencyLevel)
}
