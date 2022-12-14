package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The tail will always just move to the last point of head.
// Same principle as before. But the points are now following
// each other. So it's a list of points. The tail always follows
// the last point.
type rope struct {
	knots []point
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
	// zero is always the head.
	knots := make([]point, 0)
	for i := 0; i < 10; i++ {
		knots = append(knots, point{0, 0})
	}
	r := rope{knots: knots}

	for _, line := range split {
		var (
			dir   string
			steps int
		)
		fmt.Sscanf(line, "%s %d", &dir, &steps)
		d := directions[dir]

		for i := 0; i < steps; i++ {
			r.knots[0].x += d.x
			r.knots[0].y += d.y

			last := r.knots[0]
			for i := 1; i < len(r.knots); i++ {
				// this shouldn't be head, but the previous knot's position.
				x, y := last.x-r.knots[i].x, last.y-r.knots[i].y

				// so if both are above that means it needs to move diagonally
				// if one of them equals, we only decrease the one that is higher than one.
				if abs(x) > 1 || abs(y) > 1 {
					// if abs(x) == 0 || abs(y) == 0 {
					// we only need to move one of the positions
					if abs(x) > 1 {
						if x < 0 {
							x = -1
						} else if x > 0 {
							x = 1
						}
					}
					if abs(y) > 1 {
						if y < 0 {
							y = -1
						} else if y > 0 {
							y = 1
						}
					}
					// }
					r.knots[i].x += x
					r.knots[i].y += y
				}
				last = r.knots[i]
			}
			grid[point{x: last.x, y: last.y}]++
			// drawGrid(r.knots)
		}
	}
	// fmt.Println(grid)
	fmt.Println(len(grid))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func drawGrid(knots []point) {
	grid := [30][30]string{}
	// +15
	for i, p := range knots {
		grid[p.y+15][p.x+15] = strconv.Itoa(i)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "" {
				fmt.Print(".")
			} else {
				fmt.Print(grid[y][x])
			}
			grid[15][15] = "S"
		}
		fmt.Println()
	}
	fmt.Println("=======================")
}
