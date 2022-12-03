package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	// int('a')-int('z')+26
	// int('A')-38
	split := strings.Split(string(content), "\n")
	sum := 0
	for _, line := range split {
		head, tail := line[:len(line)/2], line[len(line)/2:]
		tailChars := make(map[rune]struct{})
		for _, c := range tail {
			tailChars[c] = struct{}{}
		}
		for _, c := range head {
			if _, ok := tailChars[c]; ok {
				// fmt.Println("found match: ", string(c))
				if unicode.IsLower(c) {
					sum += (int(c) - int('z') + 26)
				} else {
					sum += (int(c) - 38)
				}
				break
			}
		}
	}
	fmt.Println(sum)
}
