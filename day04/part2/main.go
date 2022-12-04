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

	overlaps := 0
	for _, line := range split {
		var (
			x1, x2 int
			y1, y2 int
		)
		fmt.Sscanf(line, "%d-%d,%d-%d", &x1, &x2, &y1, &y2)

		if x2 < y1 || y2 < x1 {
			continue
		}
		overlaps++
	}

	fmt.Println("overlaps: ", overlaps)
}
