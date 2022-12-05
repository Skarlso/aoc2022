package main

import (
	"fmt"
	"os"
	"strings"
)

// I could parse this into a grid and then fill up the requires slices.
// Since this will end up like a matrix anyways.
//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// [J]             [B] [W]
// [T]     [W] [F] [R] [Z]
// [Q] [M]     [J] [R] [W] [H]
// [F] [L] [P]     [R] [N] [Z] [G]
// [F] [M] [S] [Q]     [M] [P] [S] [C]
// [L] [V] [R] [V] [W] [P] [C] [P] [J]
// [M] [Z] [V] [S] [S] [V] [Q] [H] [M]
// [W] [B] [H] [F] [L] [F] [J] [V] [B]
// 1   2   3   4   5   6   7   8   9

var (
	// stacks = [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }
	stacks = [][]string{
		{"W", "M", "L", "F"},
		{"B", "Z", "V", "M", "F"},
		{"H", "V", "R", "S", "L", "Q"},
		{"F", "S", "V", "Q", "P", "M", "T", "J"},
		{"L", "S", "W"},
		{"F", "V", "P", "M", "R", "J", "W"},
		{"J", "Q", "C", "P", "N", "R", "F"},
		{"V", "H", "P", "S", "Z", "W", "R", "B"},
		{"B", "M", "J", "C", "G", "H", "Z", "W"},
	}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	for _, line := range split {
		//move 1 from 2 to 1
		var (
			move, from, to int
		)
		fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)

		var m []string
		m, stacks[from-1] = stacks[from-1][len(stacks[from-1])-move:], stacks[from-1][:len(stacks[from-1])-move]

		stacks[to-1] = append(stacks[to-1], m...)
	}

	topCratesAfterMove := ""
	for _, s := range stacks {
		topCratesAfterMove += s[len(s)-1]
	}
	fmt.Println(topCratesAfterMove)
}
