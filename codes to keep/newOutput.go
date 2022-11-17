package main

import (
	"fmt"
	"os"
	"strings"
)

var fontfile string
var n = "\n"

func main() {
	args := os.Args
	argsLen := len(args)

	flag := strings.Split(os.Args[1], "=")

	if argsLen <= 2 {
		return
	} else if argsLen > 2 {
		if flag[0] == "--output" {
			switch args[3] {
			case "shadow.txt":
				fontfile = "shadow.txt"
			case "thinkertoy":
				fontfile = "thinkertoy.txt"
				n = "\r\n"
			default:
				fontfile = "standard.txt"
			}
		}
	}
	text, _ := os.ReadFile(fontfile)
	_, err := os.Create(flag[1])
	if err != nil {
		panic(err)
	}
	s := ""

	// ascii art code
	lines := strings.Split(string(text), n)
	if argsLen <= 2 {
		fmt.Println("2 or fewer")
		split := strings.Split(os.Args[1], n)
		for i := 0; i < len(split); i++ {
			if string(split[i]) == "" {
				fmt.Println()
			} else {
				for j := 0; j < 8; j++ {
					for k := 0; k < len(string(split[i])); k++ {
						fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
						s += lines[int(((rune(split[i][k])-32)*9+1))+j]
					}
					fmt.Print(n)
					s += n
				}
				err := os.WriteFile(flag[0], []byte(s), 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
