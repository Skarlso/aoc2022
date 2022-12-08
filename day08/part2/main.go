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

	max := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			startingPoint := point{y: i, x: j}

			if score := score(startingPoint, forest); score > max {
				max = score
			}
		}
	}

	fmt.Println("max score: ", max)
}

func score(sp point, forest [][]int) int {
	score := 1
	for _, d := range directions {
		p := sp

		p.x += d.x
		p.y += d.y
		current := 0
		for {
			if p.x < 0 || p.y < 0 || p.y == len(forest) || p.x == len(forest[p.y]) {
				break
			}

			if forest[p.y][p.x] >= forest[sp.y][sp.x] {
				current++
				break
			}
			current++
			p.x += d.x
			p.y += d.y
		}
		score *= current
	}

	return score
}
