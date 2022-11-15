package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readFile() string {
	//for standard, when --output is not specified, terminal prints args[3] (the font name) instead of string
	//when font is anything else, an error occurs as there is no args[3] when output is not specified
	//therefore we need to make separate code for when there are 3 args vs when there are only 2
	if string(os.Args[2]) == "standard" || string(os.Args[3]) == "standard" {
		text, err := os.ReadFile("standard.txt")
		if err != nil {
			panic(err)
		}
		return string(text)
	} else if string(os.Args[2]) == "shadow" || string(os.Args[3]) == "shadow" {
		text, err := os.ReadFile("shadow.txt")
		if err != nil {
			panic(err)
		}
		return string(text)
	} else {
		text, err := os.ReadFile("thinkertoy.txt")
		if err != nil {
			panic(err)
		}
		return string(text)
	}
}

func main() {
	lines := strings.Split(string(readFile()), "\n")

	//OUTPUT FLAG
	output := flag.String("output", "banner.txt", "output to banner.txt?") //when false, output does nothing
	flag.Parse()

	//when output is true
	if *output == "banner.txt" {
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
			err = ioutil.WriteFile("banner.txt", []byte(s), 0644)
			if err != nil {
				panic(err)
			}
		}
	} else {
		//when output is false therefore os.Args[1] = string
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
