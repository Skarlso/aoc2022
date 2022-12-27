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

// type elf struct {
// 	id             int
// 	location       point
// 	proposingOrder []check
// }

// point would probably be enough instead of elf.
type check struct {
	name string
	f    func(elf point, grid map[point]bool) (point, bool)
}

// type check func(elf point, grid map[point]bool) (point, bool)

var (
	// Adding name so I can verify that it's shuffling them properly.
	checkNorthCheck = check{
		f:    checkNorth,
		name: "north",
	}
	checkSouthCheck = check{
		f:    checkSouth,
		name: "south",
	}
	checkWestCheck = check{
		f:    checkWest,
		name: "west",
	}
	checkEastCheck = check{
		f:    checkEast,
		name: "east",
	}
)

var propositionOrder = []check{checkNorthCheck, checkSouthCheck, checkWestCheck, checkEastCheck}

// directions are: We start of facing right.
// We will -1 current position or +1 current position based on L, R.
var directions = map[string]point{
	"N": {x: 0, y: -1}, // N
	"E": {x: 1, y: 0},  // E
	"S": {x: 0, y: 1},  // S
	"W": {x: -1, y: 0}, // W

	// cross
	"NE": {x: 1, y: -1},  // NE
	"NW": {x: -1, y: -1}, // NW
	"SW": {x: -1, y: 1},  // SW
	"SE": {x: 1, y: 1},   // SE

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

	// go through first row by row and set maxX and minX then go through column by column and set minY and maxY.
	// elfid := 0
	// elfs := make([]*elf, 0)
	// elfs := 0
	for y, line := range split {
		for x, c := range line {
			if c == '#' {
				// e := elf{
				// 	id:             elfid,
				// 	location:       point{x: x, y: y},
				// 	proposingOrder: []check{checkNorth, checkSouth, checkWest, checkEast},
				// }
				grid[point{x: x, y: y}] = true
				// elfs++
				// elfs = append(elfs, &e)
			}
		}
	}
	// fmt.Println("number of elfs: ", elfs)
	round := 0
	limit := 10
	for round < limit {
		suggestions := make(map[point][]point)
		for elf := range grid {
			// fmt.Println("elf: ", elf)

			if needsToMove(elf, grid) {
				for _, p := range propositionOrder {
					// If any of them is false, THEN we need to add the first one that is true.
					if v, ok := p.f(elf, grid); ok {
						// fmt.Printf("elf '%+v' is proposing to move to: %+v\n", elf, v)
						suggestions[v] = append(suggestions[v], elf)
						// if proposition is accepted, stop
						break
					}
				}
			}
		}
		// fmt.Println("suggestions: ", suggestions)
		for k, v := range suggestions {
			if len(v) == 1 {
				elf := v[0]
				// fmt.Printf("elf: %+v is moving to point: %+v\n", elf, k)
				// update these fuckers.
				// fmt.Println("deleting elf: ", elf)
				// fmt.Println("before delete: ", grid)
				delete(grid, elf)
				// fmt.Println("after delete: ", grid)
				// fmt.Println("adding new elf: ", k)
				grid[k] = true
				// fmt.Println("after adding: ", grid)
			}
		}
		round++
		propositionOrder = append(propositionOrder[1:], propositionOrder[0])
	}

	// fmt.Println("number of elfs: ", elfid)
	fmt.Println("number of empty places: ", countEmpty(grid))
}

func needsToMove(elf point, grid map[point]bool) bool {
	for _, d := range directions {
		if grid[point{x: elf.x + d.x, y: elf.y + d.y}] {
			return true
		}
	}

	return false
}

func countEmpty(grid map[point]bool) int {
	var (
		minx, miny        = math.MaxInt, math.MaxInt
		maxx, maxy, count int
	)
	for p := range grid {
		if p.x > maxx {
			maxx = p.x
		}
		if p.x < minx {
			minx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
		if p.y < miny {
			miny = p.y
		}

	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if !grid[point{x: x, y: y}] {
				count++
			}
		}
	}

	return count
}

func display(grid map[point]bool) {
	var (
		minx, miny = math.MaxInt, math.MaxInt
		maxx, maxy int
	)
	for p := range grid {
		if p.x > maxx {
			maxx = p.x
		}
		if p.x < minx {
			minx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
		if p.y < miny {
			miny = p.y
		}

	}
	// fmt.Printf("minx: %d, maxx: %d, miny: %d, maxy: %d\n", minx, maxx, miny, maxy)
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if grid[point{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// There is some overlap and duplicating here.
func checkNorth(elf point, grid map[point]bool) (point, bool) {
	for _, p := range []point{directions["N"], directions["NE"], directions["NW"]} {
		if _, ok := grid[point{x: elf.x + p.x, y: elf.y + p.y}]; ok {
			return point{}, false
		}
	}

	return point{x: elf.x, y: elf.y - 1}, true
}

func checkSouth(elf point, grid map[point]bool) (point, bool) {
	for _, p := range []point{directions["S"], directions["SE"], directions["SW"]} {
		if _, ok := grid[point{x: elf.x + p.x, y: elf.y + p.y}]; ok {
			return point{}, false
		}
	}

	return point{x: elf.x, y: elf.y + 1}, true
}

func checkWest(elf point, grid map[point]bool) (point, bool) {
	for _, p := range []point{directions["W"], directions["NW"], directions["SW"]} {
		if _, ok := grid[point{x: elf.x + p.x, y: elf.y + p.y}]; ok {
			return point{}, false
		}
	}

	return point{x: elf.x - 1, y: elf.y}, true
}

func checkEast(elf point, grid map[point]bool) (point, bool) {
	for _, p := range []point{directions["E"], directions["NE"], directions["SE"]} {
		if _, ok := grid[point{x: elf.x + p.x, y: elf.y + p.y}]; ok {
			return point{}, false
		}
	}

	return point{x: elf.x + 1, y: elf.y}, true
}
