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
	inputFile := flag.String("i", "", "Input file to read URLs from.")
	outputDirectory := flag.String("d", cwd,
		"Directory to save downloaded files to. Default value is the current working directory.")
	concurrencyLevel := flag.Int("c", runtime.GOMAXPROCS(0),
		"Number of concurrent downloads allowed. Default value is GOMAXPROCS (nproc, if unset).")
	flag.Parse()

	if *showHelp {
		fmt.Printf("Usage: %s [OPTION]... [URL]...\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if *concurrencyLevel < 1 {
		fmt.Fprintf(os.Stderr, "Error: concurrency level can't be less than 1\n")
		os.Exit(1)
	}

	lines := flag.Args()
	if *inputFile != "" {
		lines, err = utils.ReadLines(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
	}

	urls, err := utils.StringToURL(lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing URLs: %v\n", err)
		os.Exit(1)
	}

	err = utils.DownloadURLs(urls, *concurrencyLevel, *outputDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading: %v\n", err)
		os.Exit(1)
	}
}
