package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

// directions are: We start of facing right.
// We will -1 current position or +1 current position based on L, R.
var directions = map[string]point{
	"N": {x: 0, y: -1}, // N
	"E": {x: 1, y: 0},  // E
	"S": {x: 0, y: 1},  // S
	"W": {x: -1, y: 0}, // W

	// cross
	// "NE": {x: 1, y: -1},  // NE
	// "NW": {x: -1, y: -1}, // NW
	// "SW": {x: -1, y: 1},  // SW
	// "SE": {x: 1, y: 1},   // SE

}

// blizzard defines a blizzard with current location and direction.
type blizzard struct {
	location  point
	direction point
	// mod defines the value at which this blizzard will wrap around.
	// take head of the borders.
	mod int
	// Purely for display purposes.
	char string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}

	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	var (
		start point
		end   point
	)
	blizzards := make(map[point][]blizzard)
	grid := make([][]string, 0)
	for y, line := range split {
		if y == 0 {
			n := strings.Index(line, ".")
			if n > -1 {
				start = point{x: n, y: y}
			}
		}
		if y == len(split)-1 {
			n := strings.Index(line, ".")
			if n > -1 {
				end = point{x: n, y: y}
			}
		}
		row := make([]string, 0)
		for x, c := range strings.Split(line, "") {
			insert := c
			switch c {
			case "^":
				// -1 on the mod because of the borders.
				b := blizzard{
					location:  point{x: x, y: y},
					direction: point{x: 0, y: -1}, // N
					mod:       len(split) - 1,
					char:      "^",
				}
				blizzards[b.location] = append(blizzards[b.location], b)
				insert = "."
			case ">":
				b := blizzard{
					location:  point{x: x, y: y},
					direction: point{x: 1, y: 0}, // E
					mod:       len(line) - 1,
					char:      ">",
				}
				blizzards[b.location] = append(blizzards[b.location], b)
				insert = "."
			case "v":
				b := blizzard{
					location:  point{x: x, y: y},
					direction: point{x: 0, y: 1}, // S
					mod:       len(split) - 1,
					char:      "v",
				}
				blizzards[b.location] = append(blizzards[b.location], b)
				insert = "."
			case "<":
				b := blizzard{
					location:  point{x: x, y: y},
					direction: point{x: -1, y: 0}, // W
					mod:       len(line) - 1,
					char:      "<",
				}
				blizzards[b.location] = append(blizzards[b.location], b)
				insert = "."
			}
			row = append(row, insert)
		}
		grid = append(grid, row)
	}
	display(grid, blizzards)
	fmt.Printf("start: %+v, end: %+v\n", start, end)

	currentStep := map[point]bool{start: true}
	steps := 0
	for !currentStep[end] {
		// Before Santa moves, move the blizzards.
		blizzards = moveBlizzards(blizzards)

		newStep := make(map[point]bool)
		for p := range currentStep {
			if _, ok := blizzards[p]; !ok {
				newStep[p] = true
			}

			for _, d := range directions {
				next := point{x: p.x + d.x, y: p.y + d.y}
				// Out of bounds
				if next.x < 0 || next.y < 0 || next.x >= len(grid[next.y]) || next.y >= len(grid[next.y]) {
					continue
				}
				// But we also need to check for the end coordinate which is in the border.
				if grid[next.y][next.x] == "#" {
					continue
				}

				// there is a blizzard there
				if _, ok := blizzards[next]; ok {
					continue
				}

				newStep[next] = true
			}
		}
		currentStep = newStep
		steps++
		// display(grid, blizzards)
		// If queue is empty, we have no more moves left, we'll consider the current point a valid move.
		// This is simulating waiting.
		// var temp point
		// queue := []point{current}
		// seen := map[point]bool{current: true}
		// for

		// if len(queue) == 0 {
		// 	queue = append(queue, current)
		// }
		// current, queue = queue[0], queue[1:]
		// grid[current.y][current.x] = "S"

		// if current == end {
		// 	fmt.Println("we reached the end in steps: ", steps)
		// 	os.Exit(0)
		// }

		// // I bet that this is an endless loop.
		// queue = append(queue, neighbour(current, grid, blizzards)...)

		// steps++
	}

	fmt.Println("steps it took: ", steps)
}

// Bug: it loops over the newly created items as well. I need to create a NEW list.
// We create a new list for ever move to update the keys and make sure
// all blizzards are at the right location and they don't move twice.
func moveBlizzards(blizzards map[point][]blizzard) map[point][]blizzard {
	result := make(map[point][]blizzard)
	for k := range blizzards {
		blizzard := blizzards[k]
		for _, b := range blizzard {
			// they only move into a single direction, so it's safe to increase both one of them
			// won't change.
			b.location.x = (b.location.x + b.direction.x) % b.mod
			b.location.y = (b.location.y + b.direction.y) % b.mod
			// take care of the border
			if b.location.x == 0 {
				b.location.x = 1
			}
			if b.location.y == 0 {
				b.location.y = 1
			}
			// fmt.Println("location after: ", b.location)
			result[b.location] = append(result[b.location], b)
		}
	}
	return result
}

func neighbour(current point, grid [][]string, blizzards map[point][]blizzard) []point {
	var result []point
	// consult the blizzard map and our current location in the grid
	for _, d := range directions {
		next := point{x: current.x + d.x, y: current.y + d.y}
		// Out of bounds
		if next.x < 0 || next.y < 0 || next.x >= len(grid[next.y]) || next.y >= len(grid[next.y]) {
			continue
		}
		// But we also need to check for the end coordinate which is in the border.
		if grid[next.y][next.x] == "#" {
			continue
		}

		// there is a blizzard there
		if _, ok := blizzards[next]; ok {
			continue
		}

		result = append(result, next)
	}

	return result
}

func display(grid [][]string, blizzards map[point][]blizzard) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if v, ok := blizzards[point{x: x, y: y}]; ok {
				if len(v) > 1 {
					fmt.Print(len(v))
				} else if len(v) == 1 {
					fmt.Print(v[0].char)
				}
			} else {
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println()
	}
}
