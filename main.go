package main

import (
	"fmt"
	"os"
	"strings"

	funcs "ascii-art-fs/functions"
)

const (
	hashStandard   string = "ac85e83127e49ec42487f272d9b9db8b" // "standard.txt"
	hashShadow     string = "a49d5fcb0d5c59b2e77674aa3ab8bbb1" // "shadow.txt"
	hashThinkertoy string = "86d9947457f6a41a18cb98427e314ff8" // "thinkertoy.txt"
)

func main() {
	if len(os.Args) == 2 || len(os.Args) == 3 {
		args := os.Args[1:]
		filename := "standard.txt"
		if len(args) == 2 {
			switch args[1] {
			case "shadow":
				filename = "shadow.txt"
			case "thinkertoy":
				filename = "thinkertoy.txt"
			case "standard":
				filename = "standard.txt"
			}
		}
		if funcs.GetHash(filename) == hashStandard || funcs.GetHash(filename) == hashShadow || funcs.GetHash(filename) == hashThinkertoy {
			if args[0] == "" {
				return
			}
			if args[0] == "\\n" {
				return
			}
			for _, alter := range args[0] {
				if (rune(alter) < rune(32) || rune(alter) > rune(127)) && alter != rune(10) {
					fmt.Println("ERROR: non printable character")
					return
				}
			}
			asciiLines, err := funcs.GetStrings(filename)
			if err != nil {
				fmt.Println("ERROR: cannot READ file")
				return
			}
			asciiMap := make(map[rune][]string)
			x := 1
			y := 9
			for key := 32; key < 127; key++ {
				asciiMap[rune(key)] = asciiLines[x:y]
				x = x + 9
				y = y + 9
			}
			res := ""
			text := strings.ReplaceAll(args[0], "\n", "\\n")
			arg := strings.Split(text, "\\n")
			for i, v := range arg {
				if v == "" {
					arg[i] = ""
				}
			}
			newline := funcs.ForNewLines(arg)
			for w := 0; w < len(arg); w++ {
				if newline && w == len(arg)-1 {
					break
				}
				if arg[w] != "" {
					for i := 0; i < 8; i++ {
						for _, ch := range arg[w] {
							res = res + asciiMap[ch][i]
						}
						res = res + string(rune(10))
					}
				} else if arg[w] == "" {
					res = res + string(rune(10))
				}
			}
			fmt.Print(res)
		} else {
			fmt.Println("ERROR: Wrong hash!")
			return
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
}
