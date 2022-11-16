package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//when flag, a certain message is printed. when not, default is printed
	msg := flag.String("text", "Hello there", "message to display") //msg is a pointer for our string flag
	caps := flag.Bool("c", false, "should text be uppercase")
	nums := flag.Int("n", 1, "number of times text is displayed")
	flag.Parse() //need to parse flags after defining them

	if *caps {
		*msg = strings.ToUpper(*msg) //reassigning msg to a capitalised version
	}

	for i := 0; i < *nums; i++ {
		fmt.Println(*msg)
	}

}
