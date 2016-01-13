package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var inputFileName string
	var outputFileName string
	var encoded string

	flag.StringVar(&inputFileName, "i", "", "Input file of redis commands")
	flag.StringVar(&outputFileName, "o", "", "Output file of redis protocol")
	flag.Parse()

	input, errIn := ioutil.ReadFile(inputFileName)
	if errIn == nil {
		encoded = Encode(string(input))
		fmt.Printf(encoded)
	} else {
		fmt.Fprintf(os.Stderr, "%v", errIn)
	}
}
