package main

import (
	"fmt"
	"os"
	"strings"
)

var j = len(os.Args) - 1

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

func main() {
	switch os.Args[j] {
	case "thinkertoy":
		lines := strings.Split(string(readFile()), "\r\n") //thinkertoy NEEDS a carriage return
		split := strings.Split(os.Args[1], `\r\n`)
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
	default:
		lines := strings.Split(string(readFile()), "\n")
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
