package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type direction struct {
	x, y int
}

var directions = []direction{
	{x: 0, y: +1}, //up
	{x: 0, y: -1}, //down
	{x: -1, y: 0}, //left
	{x: +1, y: 0}, //right
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")

	forest := make([][]int, 0)

	for _, line := range split {
		row := make([]int, 0)
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		forest = append(forest, row)
	}

	trees := 0
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			startingPoint := point{i, j}

			if isVisible(startingPoint, forest) {
				trees++
			}
		}
	}

	fmt.Println("total hidden trees: ", trees)
}

func isVisible(sp point, forest [][]int) bool {
	// If we find a single path which leads to the edge, we know its visible.
	// But it can only go straight.
	for _, d := range directions {
		p := sp

		for {
			if p.x+d.x == 0 || p.x+d.x >= len(forest[p.y+d.y])-1 || p.y+d.y == 0 || p.y+d.y >= len(forest)-1 {
				return false
			}
			if forest[p.y+d.y][p.x+d.x] >= forest[sp.y][sp.x] {
				break
			}

			p.x += d.x
			p.y += d.y
		}

	}

	return true
}
