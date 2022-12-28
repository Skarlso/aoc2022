package main

import (
	"fmt"
	"os"
	"strings"
)

type cube struct {
	sideHashes []string
	location   point
}

func (c cube) calculateSides() {
	// take the location and calculate the rest of the sides

}

type point struct {
	x, y, z int
}

func (p point) String() string {
	return fmt.Sprintf("%d%d%d", p.x, p.y, p.z)
}

// sides
var sides = []point{
	{x: 0, y: 0, z: 1},
	{x: 0, y: 1, z: 0},
	{x: 1, y: 0, z: 0},
	{x: 0, y: 0, z: -1},
	{x: 0, y: -1, z: 0},
	{x: -1, y: 0, z: 0},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	globalSides := make(map[point]int)

	for _, line := range split {
		if line == "" {
			continue
		}
		var (
			x, y, z int
		)
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		p := point{x: x, y: y, z: z}
		for _, s := range sides {
			globalSides[point{x: p.x + s.x, y: p.y + s.y, z: p.z + s.z}]++
		}
	}

	totalLoneSides := 0
	for _, v := range globalSides {
		if v == 1 {
			totalLoneSides++
		}
	}
	fmt.Println("lone sides: ", totalLoneSides)
}
