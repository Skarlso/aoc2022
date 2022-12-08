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
	{x: 0, y: +1}, //down
	{x: 0, y: -1}, //up
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
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			startingPoint := point{x: j, y: i}

			if isVisible(startingPoint, forest) {
				// fmt.Println("visible: ", forest[startingPoint.y][startingPoint.x])
				trees++
			}
		}
	}

	fmt.Println("total visible trees: ", trees)
}

func isVisible(sp point, forest [][]int) bool {
	// If we find a single path which leads to the edge, we know its visible.
	// But it can only go straight.
	for _, d := range directions {
		p := sp

		p.x += d.x
		p.y += d.y
		// fmt.Println("starting point: ", sp)
		for {
			// fmt.Println("point: ", p)
			if p.x < 0 || p.y < 0 || p.y == len(forest) || p.x == len(forest[p.y]) {
				// fmt.Println("reached end")
				return true
			}

			// fmt.Printf("comparing origin '%d' with '%d'\n", forest[sp.y][sp.x], forest[p.y][p.x])
			if forest[p.y][p.x] >= forest[sp.y][sp.x] {
				// fmt.Println("break")
				break
			}

			p.x += d.x
			p.y += d.y
		}

	}

	return false
}
