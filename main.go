package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var inputFileName string

	flag.StringVar(&inputFileName, "i", "", "Input file of redis commands")
	flag.Parse()

	file, err := os.Open(inputFileName)
	if err == nil {
		EncodeStream(file, os.Stdout)
	} else {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
