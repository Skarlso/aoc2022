package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	grid := make(map[point]bool)
	maxy := 0
	maxx := 0
	minx := math.MaxInt
	for _, line := range split {
		borders := strings.Split(line, " -> ")
		for i := 0; i < len(borders)-1; i++ {
			var (
				sFromX, sFromY int
				sToX, sToY     int
			)
			fmt.Sscanf(borders[i], "%d,%d", &sFromX, &sFromY)
			fmt.Sscanf(borders[i+1], "%d,%d", &sToX, &sToY)
			if sToY > maxy {
				maxy = sToY
			}
			if sToX > maxx {
				maxx = sToX
			}
			if sFromX > maxx {
				maxx = sFromX
			}
			if sFromX < minx {
				minx = sFromX
			}
			if sToX < minx {
				minx = sToX
			}

			var (
				fromx, fromy, tox, toy int
			)
			if sFromX > sToX {
				tox = sFromX
				fromx = sToX
			} else {
				fromx = sFromX
				tox = sToX
			}
			if sFromY > sToY {
				fromy = sToY
				toy = sFromY
			} else {
				fromy = sFromY
				toy = sToY
			}
			for x := fromx; x <= tox; x++ {
				for y := fromy; y <= toy; y++ {
					grid[point{x: x, y: y}] = true
				}
			}
		}
	}

	maxy += 2
	draw(minx, maxx, maxy, grid)
	// A grain is falling until it reached the maximum y coordinate.
	count := 0
	start := &point{x: 500, y: 0}
	for {
		// I don't think we need a _visited_ because we'll be going only DOWN and left and right if there is space.
		// if we bump into a location that has an item in it and there is nowhere to go, then we stop the loop anyways.
		queue := []*point{start}
		var current *point

		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]
			next := falling(current, grid, maxy)
			// There is nowhere to go
			if next == nil {
				// If we just started but there is nowhere for the grain to fall to
				// we can assume that the grid is full.
				if current == start {
					fmt.Println("maximum depth: ", maxy)
					// plus one for the starting point
					fmt.Println("number of sand grains: ", count+1)
					os.Exit(0)
				}
				// save the current location as the grain coming to rest.
				grid[*current] = true
				break
			}
			queue = append(queue, next)
		}

		count++
	}
}

func falling(p *point, grid map[point]bool, maxy int) *point {
	if p.y+1 == maxy {
		return nil
	}
	if !grid[point{x: p.x, y: p.y + 1}] {
		return &point{x: p.x, y: p.y + 1}
	} else {
		if !grid[point{x: p.x - 1, y: p.y + 1}] {
			return &point{x: p.x - 1, y: p.y + 1}
		} else if !grid[point{x: p.x + 1, y: p.y + 1}] {
			return &point{x: p.x + 1, y: p.y + 1}
		}
	}
	return nil
}

func draw(fromx, maxx, maxy int, grid map[point]bool) {
	for y := 0; y <= maxy; y++ {
		for x := fromx; x <= maxx; x++ {
			if grid[point{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
