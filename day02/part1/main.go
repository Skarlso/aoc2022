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
			"X": 1,
			"Y": 2,
			"Z": 3,
		}
		dic = map[string]string{
			"A": "R",
			"X": "R",
			"B": "P",
			"Y": "P",
			"C": "S",
			"Z": "S",
		}
	)

	sum := 0
	for _, l := range split {
		split := strings.Split(l, " ")
		op := split[0]
		me := split[1]

		if dic[op] == dic[me] {
			sum += mineScore[me] + 3
		} else if dic[op] == "R" {
			if dic[me] == "S" {
				sum += mineScore[me] + 0
			} else if dic[me] == "P" {
				sum += mineScore[me] + 6
			}
		} else if dic[op] == "P" {
			if dic[me] == "S" {
				sum += mineScore[me] + 6
			} else if dic[me] == "R" {
				sum += mineScore[me] + 0
			}
		} else if dic[op] == "S" {
			if dic[me] == "R" {
				sum += mineScore[me] + 6
			} else if dic[me] == "P" {
				sum += mineScore[me] + 0
			}
		}
	}
	fmt.Println("score: ", sum)
}
