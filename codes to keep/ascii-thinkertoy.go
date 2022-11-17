package ascii

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func asciiToy(s string) {
	wordArg := os.Args
	split := strings.Split(wordArg[j], `\n`)
	text, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(text), "\r\n")
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
