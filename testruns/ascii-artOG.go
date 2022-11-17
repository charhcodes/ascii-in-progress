package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	wordArg := os.Args
	split := strings.Split(wordArg[1], `\n`)
	text, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(string(split[i])); k++ {
					fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
				}
				fmt.Println()
			}
		}
	}
}
