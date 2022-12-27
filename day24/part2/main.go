package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

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
	// the point to wrap too
	wrappingPoint point
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
	var blizzards []*blizzard
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
					location:      point{x: x, y: y},
					direction:     point{x: 0, y: -1}, // N
					wrappingPoint: point{x: x, y: len(split) - 2},
					char:          "^",
				}
				blizzards = append(blizzards, &b)
				insert = "."
			case ">":
				b := blizzard{
					location:      point{x: x, y: y},
					direction:     point{x: 1, y: 0}, // E
					wrappingPoint: point{x: 1, y: y},
					char:          ">",
				}
				blizzards = append(blizzards, &b)
				insert = "."
			case "v":
				b := blizzard{
					location:      point{x: x, y: y},
					direction:     point{x: 0, y: 1}, // S
					wrappingPoint: point{x: x, y: 1},
					char:          "v",
				}
				blizzards = append(blizzards, &b)
				insert = "."
			case "<":
				b := blizzard{
					location:      point{x: x, y: y},
					direction:     point{x: -1, y: 0}, // W
					wrappingPoint: point{x: len(line) - 2, y: y},
					char:          "<",
				}
				blizzards = append(blizzards, &b)
				insert = "."
			}
			row = append(row, insert)
		}
		grid = append(grid, row)
	}
	fmt.Printf("start: %+v, end: %+v\n", start, end)
	bm := make(map[point]blizzard)
	for _, b := range blizzards {
		bm[b.location] = *b
	}
	// fmt.Println("initial state: ")
	// display(grid, bm)
	steps := dodgeBlizzards(start, end, grid, blizzards)
	fmt.Println("done to end")
	steps += dodgeBlizzards(end, start, grid, blizzards)
	fmt.Println("done to start")
	steps += dodgeBlizzards(start, end, grid, blizzards)
	fmt.Println("done to back to end")
	fmt.Println("steps it took: ", steps)
}

func dodgeBlizzards(start, end point, grid [][]string, blizzards []*blizzard) int {
	currentStep := map[point]bool{start: true}
	steps := 0
	for !currentStep[end] {
		// Before Santa moves, move the blizzards.
		blizzardMap := moveBlizzards(grid, blizzards)
		// display(grid, blizzardMap)
		newStep := make(map[point]bool)
		for p := range currentStep {
			if _, ok := blizzardMap[p]; !ok {
				newStep[p] = true
			}

			for _, d := range directions {
				next := point{x: p.x + d.x, y: p.y + d.y}
				_, ok := blizzardMap[next]
				if next.x >= 0 && next.y >= 0 && next.y < len(grid) && next.x < len(grid[next.y]) && grid[next.y][next.x] != "#" && !ok {
					newStep[next] = true
				}
			}
		}
		currentStep = newStep
		steps++
	}

	return steps
}

func moveBlizzards(grid [][]string, blizzards []*blizzard) map[point]blizzard {
	result := make(map[point]blizzard)
	for b := 0; b < len(blizzards); b++ {
		// they only move into a single direction, so it's safe to increase both one of them
		// won't change.
		blizzards[b].location.x = blizzards[b].location.x + blizzards[b].direction.x
		blizzards[b].location.y = blizzards[b].location.y + blizzards[b].direction.y

		if grid[blizzards[b].location.y][blizzards[b].location.x] == "#" {
			blizzards[b].location = blizzards[b].wrappingPoint
		}

		result[blizzards[b].location] = *blizzards[b]
	}
	return result
}

func display(grid [][]string, blizzards map[point]blizzard) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if v, ok := blizzards[point{x: x, y: y}]; ok {
				fmt.Print(v.char)
			} else {
				fmt.Print(grid[y][x])
			}
		}
		fmt.Println()
	}
}
