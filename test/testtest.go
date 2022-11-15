

// case "thinkertoy":
// 	lines := strings.Split(string(readFile()), "\r\n") //thinkertoy NEEDS a carriage return
// 	if *output == "banner.txt" { //when thinkertoy && output == false, strings input will be os.Args[1]
// 		split := strings.Split(os.Args[1], `\r\n`)
// 		for i := 0; i < len(split); i++ {
// 			if string(split[i]) == "" {
// 				fmt.Println()
// 			} else {
// 				for j := 0; j < 8; j++ {
// 					for k := 0; k < len(string(split[i])); k++ {
// 						fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
// 					}
// 					fmt.Print("\r\n")
// 				}
// 			}
// 		}
// 	} else { //when thinkertoy && output == true, string input is os.Args[2]
// 		split := strings.Split(os.Args[2], `\r\n`)
// 		//create banner.txt file
// 		_, err := os.Create("banner.txt")
// 		if err != nil {
// 			panic(err)
// 		}
// 		s := ""
// 		for i := 0; i < len(split); i++ {
// 			if string(split[i]) == "" {
// 				fmt.Println()
// 			} else {
// 				for j := 0; j < 8; j++ {
// 					for k := 0; k < len(string(split[i])); k++ {
// 						fmt.Print(lines[int(((rune(split[i][k])-32)*9+1))+j])
// 						s += lines[int(((rune(split[i][k])-32)*9+1))+j]
// 					}
// 					fmt.Print("\r\n")
// 					s += "\r\n"
// 				}
// 			}
// 			err = os.WriteFile("banner.txt", []byte(s), 0644)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}


