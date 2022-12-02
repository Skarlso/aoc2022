package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	var (
		mineScore = map[string]int{
			"R": 1,
			"P": 2,
			"S": 3,
		}
		dic = map[string]string{
			"A": "R",
			"B": "P",
			"C": "S",
		}
		outcome = map[string]map[string]string{
			"X": {
				"R": "S",
				"P": "R",
				"S": "P",
			},
			"Y": {
				"R": "R",
				"P": "P",
				"S": "S",
			},
			"Z": {
				"R": "P",
				"P": "S",
				"S": "R",
			},
		}
		win = map[string]int{
			"X": 0,
			"Y": 3,
			"Z": 6,
		}
	)

	sum := 0
	for _, l := range split {
		split := strings.Split(l, " ")
		op := split[0]
		me := split[1]
		sum += mineScore[outcome[me][dic[op]]] + win[me]
	}
	fmt.Println("score: ", sum)
}
