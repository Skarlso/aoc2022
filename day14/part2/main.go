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
	count := 0
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

	draw(minx, maxx, maxy, grid)
	fmt.Println("maximum depth: ", maxy)
	fmt.Println("number of sand grains: ", count)
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
