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
	var group []string

	// read by line of threes.
	for len(split) > 0 {
		group, split = split[0:3], split[3:]
		head := group[0]
		middle := group[1]
		tail := group[2]

		headChars := make(map[rune]struct{})
		middleChars := make(map[rune]struct{})
		tailChars := make(map[rune]struct{})
		for _, c := range head {
			headChars[c] = struct{}{}
		}
		for _, c := range middle {
			middleChars[c] = struct{}{}
		}
		for _, c := range tail {
			tailChars[c] = struct{}{}
		}

		groupChars := make(map[rune]int)

		for k := range headChars {
			groupChars[k]++
		}
		for k := range middleChars {
			groupChars[k]++
		}
		for k := range tailChars {
			groupChars[k]++
		}
		for k, v := range groupChars {
			if v == 3 {
				if unicode.IsLower(k) {
					sum += (int(k) - int('z') + 26)
				} else {
					sum += (int(k) - 38)
				}
			}
		}
	}
	fmt.Println(sum)
}
