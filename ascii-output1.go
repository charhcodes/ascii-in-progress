package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var j = len(os.Args) - 1

// OUTPUT FLAG
var output = flag.String("output", "banner.txt", "output to banner.txt?") //when false, output does nothing

// readFile checks os.Args for the font
func readFile() string {
	if string(os.Args[j]) == "thinkertoy" {
		text, _ := os.ReadFile("thinkertoy.txt")
		return string(text)
	} else if string(os.Args[j]) == "shadow" {
		text, _ := os.ReadFile("shadow.txt")
		return string(text)
	} else {
		text, _ := os.ReadFile("standard.txt")
		return string(text)
	}
}

// code ignores checks to the output flag
// hence code sees our string input as os.Args[2] for all terminal arguments
// how to check whether os.Args[1] is a flag?
func outputChecker() bool {
	checker := false
	for _, arg := range os.Args {
		if arg == `--output="banner.txt"` {
			checker = true
			fmt.Println("checker working")
		}
	}
	return checker
}

func main() {
	flag.Parse()
	switch os.Args[j] {
	//THINKERTOY
	case "thinkertoy":
		lines := strings.Split(string(readFile()), "\r\n") //thinkertoy NEEDS a carriage return
		if outputChecker() {
			fmt.Println("output=true") //when thinkertoy && output == true, strings input will be os.Args[2]
			split := strings.Split(os.Args[2], `\r\n`)
			for i := 0; i < len(split); i++ {
				if string(split[i]) == "" {
					fmt.Println()
				} else {
					for j := 0; j < 8; j++ {
						for k := 0; k < len(string(split[i])); k++ {
							fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
						}
						fmt.Print("\r\n")
					}
				}
			}
		} else if !outputChecker() { //when thinkertoy && output == false, string input is os.Args[1]
			//THIS ALWAYS RUNS REGARDLESS OF TERMINAL INPUT
			fmt.Println("output=false")
			split := strings.Split(os.Args[1], `\r\n`)
			//create banner.txt file
			_, err := os.Create("banner.txt")
			if err != nil {
				panic(err)
			}
			s := ""
			for i := 0; i < len(split); i++ {
				if string(split[i]) == "" {
					fmt.Println()
				} else {
					for j := 0; j < 8; j++ {
						for k := 0; k < len(string(split[i])); k++ {
							fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
							s += lines[int(((rune(split[i][k])-32)*9+1))+j]
						}
						fmt.Print("\r\n")
						s += "\r\n"
					}
				}
				err = os.WriteFile("banner.txt", []byte(s), 0644)
				if err != nil {
					panic(err)
				}
			}
		}

	// shadow and standard still require editing, will return to default once thinkertoy is sorted out
	//STANDARD & SHADOW
	default:
		lines := strings.Split(string(readFile()), "\n")
		switch *output == "banner.txt" { //when output == true
		case true:
			split := strings.Split(os.Args[2], `\n`)
			//create banner.txt file
			_, err := os.Create("banner.txt")
			if err != nil {
				panic(err)
			}
			s := ""
			for i := 0; i < len(split); i++ {
				if string(split[i]) == "" {
					fmt.Println()
				} else {
					for j := 0; j < 8; j++ {
						for k := 0; k < len(string(split[i])); k++ {
							fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
							s += lines[int(((rune(split[i][k])-32)*9+1))+j]
						}
						fmt.Print("\n")
						s += "\n"
					}
				}
				err = os.WriteFile("banner.txt", []byte(s), 0644)
				if err != nil {
					panic(err)
				}
			}
		//when output == false
		default:
			split := strings.Split(os.Args[1], `\n`)
			for i := 0; i < len(split); i++ {
				if string(split[i]) == "" {
					fmt.Println()
				} else {
					for j := 0; j < 8; j++ {
						for k := 0; k < len(string(split[i])); k++ {
							fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
						}
						fmt.Print("\n")
					}
				}
			}
		}
	}
}
