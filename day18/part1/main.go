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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	globalSides := make(map[string]int)
	for _, line := range split {
		if line == "" {
			continue
		}
		var (
			x, y, z int
		)
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

		// up
		up := fmt.Sprintf("%s%s%s%s", point{x: x, y: y + 1, z: z}, point{x: x, y: y + 1, z: z + 1}, point{x: x + 1, y: y + 1, z: z + 1}, point{x: x + 1, y: y + 1, z: z})
		globalSides[up]++
		// down
		down := fmt.Sprintf("%s%s%s%s", point{x: x, y: y, z: z}, point{x: x, y: y, z: z + 1}, point{x: x + 1, y: y, z: z + 1}, point{x: x + 1, y: y, z: z})
		globalSides[down]++
		// left
		left := fmt.Sprintf("%s%s%s%s", point{x: x, y: y, z: z}, point{x: x, y: y + 1, z: z}, point{x: x, y: y + 1, z: z + 1}, point{x: x, y: y, z: z + 1})
		globalSides[left]++
		// right
		right := fmt.Sprintf("%s%s%s%s", point{x: x + 1, y: y, z: z}, point{x: x + 1, y: y + 1, z: z}, point{x: x + 1, y: y + 1, z: z + 1}, point{x: x + 1, y: y, z: z + 1})
		globalSides[right]++
		// front
		front := fmt.Sprintf("%s%s%s%s", point{x: x, y: y, z: z}, point{x: x, y: y + 1, z: z}, point{x: x + 1, y: y, z: z}, point{x: x + 1, y: y + 1, z: z})
		globalSides[front]++
		// back
		back := fmt.Sprintf("%s%s%s%s", point{x: x, y: y, z: z + 1}, point{x: x, y: y + 1, z: z + 1}, point{x: x + 1, y: y, z: z + 1}, point{x: x + 1, y: y + 1, z: z + 1})
		globalSides[back]++

		//hash := fmt.Sprintf("%s-%s-%s-%s-%s-%s", up, down, left, right, front, back)
		// Create a hash of all sides and put them into a global map[string]int.
		// Anything with 1 is a side that doesn't have a pair.
		//globalSides[hash]++
	}
	loneSides := 0
	for _, v := range globalSides {
		if v == 1 {
			loneSides++
		}
	}
	fmt.Println("lone sides: ", loneSides)
}
