package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := ioutil.ReadFile(file)

	split := strings.Split(string(content), "\n")
	max := 0
	sumSoFar := 0
	for _, l := range split {
		if l == "" {
			// todo: next elf
			if sumSoFar > max {
				max = sumSoFar
			}
			sumSoFar = 0
			continue
		}
		i, _ := strconv.Atoi(strings.Trim(l, "\n"))
		sumSoFar += i
	}
	fmt.Println("fattest elf: ", max)
}
