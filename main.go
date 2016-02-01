package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var inputFileName string
	var outputFileName string
	var outputWriter io.Writer

	flag.StringVar(&inputFileName, "i", "", "Input file of redis commands")
	flag.StringVar(&outputFileName, "o", "", "Output file in redis protocol format")
	flag.Parse()

	inputFile, ierr := os.Open(inputFileName)
	outputFile, oerr := os.Create(outputFileName)
	defer inputFile.Close()
	defer outputFile.Close()

	if ierr == nil {
		if oerr == nil {
			outputWriter = outputFile
		} else {
			outputWriter = os.Stdout
		}
		bufwriter := bufio.NewWriterSize(outputWriter, 4294967296)
		EncodeStream(inputFile, bufwriter)
		bufwriter.Flush()
	} else {
		fmt.Fprintf(os.Stderr, "input error: %\noutput error: %v", ierr, oerr)
	}
}
