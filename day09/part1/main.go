package main

import (
	"fmt"
	"os"
	"strings"
)

// The tail will always just move to the last point of head.
type rope struct {
	headPosition point
	tailPosition point
}

type point struct {
	x, y int
}

var directions = map[string]point{
	"D": {x: 0, y: +1}, //down
	"U": {x: 0, y: -1}, //up
	"L": {x: -1, y: 0}, //left
	"R": {x: +1, y: 0}, //right
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")

	// grid tracks the point movement of the TAIL
	// it will count up every time a point is seen.
	// the length of the grid will tell how many unique points
	// the tail was in.
	grid := make(map[point]int)
	r := rope{headPosition: point{0, 0}, tailPosition: point{0, 0}}

	for _, line := range split {
		var (
			dir   string
			steps int
		)
		fmt.Sscanf(line, "%s %d", &dir, &steps)
		d := directions[dir]
		// i will track how many steps to take.
		for i := 0; i < steps; i++ {
			newHeadX := r.headPosition.x + d.x
			newHeadY := r.headPosition.y + d.y

			x, y := newHeadX-r.tailPosition.x, newHeadY-r.tailPosition.y
			if abs(x) > 1 || abs(y) > 1 {
				r.tailPosition.x = r.headPosition.x
				r.tailPosition.y = r.headPosition.y
			}
			r.headPosition.x = newHeadX
			r.headPosition.y = newHeadY

			grid[point{x: r.tailPosition.x, y: r.tailPosition.y}]++
		}
	}

	fmt.Println(len(grid))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
