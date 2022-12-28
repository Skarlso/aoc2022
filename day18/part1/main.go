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

	globalSides := make(map[point]bool)

	listOfPoints := make([]point, 0)
	for _, line := range split {
		if line == "" {
			continue
		}
		var (
			x, y, z int
		)
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		p := point{x: x, y: y, z: z}
		listOfPoints = append(listOfPoints, p)
		globalSides[p] = true
	}

	totalLoneSides := 0
	for _, p := range listOfPoints {
		loneSides := 6
		for _, s := range sides {
			if globalSides[point{x: p.x + s.x, y: p.y + s.y, z: p.z + s.z}] {
				loneSides--
			}
		}
		totalLoneSides += loneSides
	}
	fmt.Println("lone sides: ", totalLoneSides)
}
