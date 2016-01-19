package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"io/ioutil"
	"os"
)

func main() {
	var inputFileName string
	var outputFileName string
	var encoded string

	defer profile.Start(profile.MemProfile).Stop()

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
