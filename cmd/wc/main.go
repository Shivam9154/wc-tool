package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"wc-tool/internal/counter"
	"wc-tool/pkg/utils"
)

func main() {
	countBytes := flag.Bool("c", false, "count bytes")
	countLines := flag.Bool("l", false, "count lines")
	countWords := flag.Bool("w", false, "count words")
	countChars := flag.Bool("m", false, "count characters")
	flag.Parse()

	if !(*countBytes || *countLines || *countWords || *countChars) {
		*countBytes = true
		*countLines = true
		*countWords = true
		*countChars = true
	}

	var input io.Reader
	var filename string

	args := flag.Args()
	if len(args) > 0 {
		filename = args[0]
		file, err := os.Open(filename)
		utils.CheckError(err, "Error Opening File")
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	c := counter.New()
	err := c.Process(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing input: %v\n", err)
		os.Exit(1)
	}

	var results []interface{}
	var format string

	if *countLines {
		results = append(results, c.Lines())
		format += "%8d"
	}
	if *countWords {
		results = append(results, c.Words())
		format += "%8d"
	}
	if *countBytes {
		results = append(results, c.Bytes())
		format += "%8d"
	}
	if *countChars {
		results = append(results, c.Characters())
		format += "%8d"
	}
	if filename != "" {
		results = append(results, filename)
		format += "%s"
	}

	format += "\n"
	fmt.Printf(format, results...)
}
