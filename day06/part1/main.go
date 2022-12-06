package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	checkForDuplicates := func(list []byte) bool {
		set := make(map[byte]int)
		for _, c := range list {
			set[c]++
			if set[c] > 1 {
				return true
			}
		}
		return false
	}
	for i := range content {
		if !checkForDuplicates(content[i : i+4]) {
			fmt.Println("char: ", i+4)
			break
		}
	}
}
