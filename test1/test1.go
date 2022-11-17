package main

import (
	"flag"
	"os"
	"asciiart"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	output := flag.String("output", "banner.txt", "input file name here")
	converter()
}
