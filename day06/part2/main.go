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

	unique := func(list []byte) bool {
		set := make(map[byte]int)
		for _, c := range list {
			set[c]++
			if set[c] > 1 {
				return false
			}
		}
		return true
	}
	for i := range content {
		if unique(content[i : i+14]) {
			fmt.Println("char: ", i+14)
			break
		}
	}
}
