package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// if the row has a -1 miny or maxy that means we don't care about it.
// if it has a none negative y, we check for wrapping as we hit the border.
type row struct {
	minX, maxX, minY, maxY int
	items                  []string
}

type santa struct {
	direction int
	position  point
}

type point struct {
	x, y int
}

// directions are: We start of facing right.
// We will -1 current position or +1 current position based on L, R.
var directions = []point{
	{x: 1, y: 0},  // facing right
	{x: 0, y: 1},  // facing down
	{x: -1, y: 0}, // facing left
	{x: 0, y: -1}, // facing up
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	readPath := false
	path := ""
	rows := make([]row, 0)
	// I need to find a way to find where a tile begins. I guess if the next max/min X changes or something.
	// then we know we are at a new tile. We need the max and min Y because it could wrap around there as well.
	for y, line := range split {
		// We must set the previous rows maxY here.
		if line == "" {
			rows[y-1].maxY = y - 1
			readPath = true
			continue
		}
		if readPath {
			path = line
			continue
		}
		r := row{
			items: strings.Split(line, ""),
			minX:  math.MaxInt,
		}

		if v := strings.LastIndex(line, "#"); v > -1 {
			if v > r.maxX {
				r.maxX = v
			}
		}
		if v := strings.LastIndex(line, "."); v > -1 {
			if v > r.maxX {
				r.maxX = v
			}
		}
		if v := strings.Index(line, "#"); v > -1 {
			if v < r.minX {
				r.minX = v
			}
		}
		if v := strings.Index(line, "."); v > -1 {
			if v < r.minX {
				r.minX = v
			}
		}

		// We set all rows to 0 for Y except for the borders so we can modulo freely without
		// messing up the down or up move.
		if y == 0 {
			r.minY = 1
			r.maxY = 1
		} else {
			r.minY = 1
			r.maxY = 1
		}
		// if there is a row before this row
		if y-1 > 0 {
			// if one of the x values doesn't match this row, we know we are in a new tile.
			if rows[y-1].maxX != r.maxX || rows[y-1].minX != r.minX {
				rows[y-1].maxY = y - 1
				r.minY = y
			}
		}

		rows = append(rows, r)
	}

	s := santa{
		position:  point{x: rows[0].minX, y: 0},
		direction: 0,
	}

	fmt.Println("santa starting position: x, y", s.position.x, s.position.y)
	// collect until a letter appears
	number := ""
	for i, c := range path {
		// move as much as is collected including wrapping, then
		// get the next char and rotate. repeat until done.
		// there is a single move at the end of the path. Deal with that.
		if unicode.IsDigit(c) && i < len(path)-1 {
			number += string(c)
		} else if unicode.IsLetter(c) {
			// move
			n, _ := strconv.Atoi(number)

			for m := 0; m < n; m++ {
				// if next position that we are going to move to is a wall, break.
				currentRow := rows[s.position.y]
				// +1 to allow to check if the next move would be a wall item.
				x := mod(s.position.x+directions[s.direction].x, currentRow.maxX+1)
				if x == 0 && s.position.x != 0 { // only change it if we weren't at 0 to begin with
					x += currentRow.minX
				}
				y := mod(s.position.y+directions[s.direction].y, currentRow.maxY+1)
				if y == 0 && s.position.y != 0 {
					y += currentRow.minY
				}
				fmt.Printf("moving to x: %d; y: %d\n", x, y)
				if rows[y].items[x] == "#" {
					fmt.Println("hit wall")
					break
				}
				s.position.x = x
				s.position.y = y
				switch s.direction {
				case 0:
					rows[s.position.y].items[s.position.x] = ">"
				case 1:
					rows[s.position.y].items[s.position.x] = "v"
				case 2:
					rows[s.position.y].items[s.position.x] = "<"
				case 3:
					rows[s.position.y].items[s.position.x] = "^"
				}
				// display(s, rows)
				// next := point{
				// plus minX because the smallest is not 0.
				// x: mod(s.position.x+s.direction.x, currentRow.maxX) + currentRow.minX,
				// }
				// var y int
			}

			// rotate in direction c
			if c == 'L' {
				fmt.Println("rotating left from: ", s.direction)
				s.direction = mod(s.direction-1, len(directions))
				fmt.Println("rotating left to: ", s.direction)
			} else if c == 'R' {
				fmt.Println("rotating right from: ", s.direction)
				s.direction = mod(s.direction+1, len(directions))
				fmt.Println("rotating right to: ", s.direction)
			}

			// clear number, start again
			number = ""
		} else {
			// final move
			number += string(c)
			n, _ := strconv.Atoi(number)
			for m := 0; m < n; m++ {
				// if next position that we are going to move to is a wall, break.
				currentRow := rows[s.position.y]
				x := mod(s.position.x+directions[s.direction].x, currentRow.maxX)
				if x == 0 && s.position.x != 0 { // only change it if we weren't at 0 to begin with
					x += currentRow.minX
				}
				y := mod(s.position.y+directions[s.direction].y, currentRow.maxY)
				if y == 0 && s.position.y != 0 {
					y += currentRow.minY
				}
				// fmt.Printf("final move to x: %d; y: %d\n", x, y)
				if rows[y].items[x] == "#" {
					fmt.Println("hit wall")
					break
				}
				s.position.x = x
				s.position.y = y
				switch s.direction {
				case 0:
					rows[s.position.y].items[s.position.x] = ">"
				case 1:
					rows[s.position.y].items[s.position.x] = "v"
				case 2:
					rows[s.position.y].items[s.position.x] = "<"
				case 3:
					rows[s.position.y].items[s.position.x] = "^"
				}
				// display(s, rows)
				// next := point{
				// plus minX because the smallest is not 0.
				// x: mod(s.position.x+s.direction.x, currentRow.maxX) + currentRow.minX,
				// }
				// var y int
			}
			// fmt.Println("final move n: ", n)
		}

	}

	// If we move into a wall we `break` as we can't go forward any longer.
	// fmt.Println("starting position: ", s)
	// fmt.Printf("rows: %+v", rows)
	// fmt.Println("path: ", path)
	display(rows)
	fmt.Printf("row: %d; column: %d; facing: %d \n", s.position.y+1, s.position.x+1, s.direction)

}

func display(rows []row) {
	for _, r := range rows {
		for x := 0; x <= r.maxX; x++ {

			fmt.Print(string(r.items[x]))
		}
		fmt.Println()
	}
	// for _, r := range rows {
	// 	for x := r.minX; x <= r.maxX; x++ {
	// 		fmt.Print(string(r.items[x]))
	// 	}
	// 	fmt.Println()
	// }
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
