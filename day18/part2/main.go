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

	for _, line := range split {
		if line == "" {
			continue
		}
		var (
			x, y, z int
		)
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		p := point{x: x, y: y, z: z}
		globalSides[p] = true
	}

	var (
		maxx, maxy, maxz int
	)
	for k := range globalSides {
		if k.x > maxx {
			maxx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
		if k.z > maxz {
			maxz = k.z
		}
	}

	totalEdgeFacingSides := 0
	for p := range globalSides {
		for _, s := range sides {
			next := point{
				x: p.x + s.x,
				y: p.y + s.y,
				z: p.z + s.z,
			}

			if reachesEdge(next, globalSides, maxx, maxy, maxz) {
				totalEdgeFacingSides++
			}
		}
	}

	fmt.Println("total edge facing sides: ", totalEdgeFacingSides)
}

func reachesEdge(p point, totalSides map[point]bool, maxx, maxy, maxz int) bool {
	queue := []point{p}
	seen := map[point]bool{}
	var current point
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if seen[current] || totalSides[current] {
			continue
		}
		seen[current] = true

		// edge reached
		if current.x <= 0 || current.x >= maxx ||
			current.y <= 0 || current.y >= maxy ||
			current.z <= 0 || current.z >= maxz {
			return true
		}

		for _, s := range sides {
			next := point{
				x: current.x + s.x,
				y: current.y + s.y,
				z: current.z + s.z,
			}
			queue = append(queue, next)
		}
	}
	return false
}
