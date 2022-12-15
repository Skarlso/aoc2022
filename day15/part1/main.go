package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run part1/main.go [file] [test-row-number]")
		os.Exit(1)
	}
	file := os.Args[1]
	testRow := os.Args[2]
	row, err := strconv.Atoi(testRow)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	// This could just be bool but I'm using string so I can visually verify that things look okay.
	// End this helped me a lot because I saw that I needed distance <= maxDistance instead of distance < maxDistance.
	grid := make(map[point]string)
	var (
		minx, miny = math.MaxInt, math.MaxInt
		maxx, maxy int
	)
	for _, line := range split {
		var (
			sensorx, sensory int
			beaconx, beacony int
		)
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorx, &sensory, &beaconx, &beacony)
		sensor := point{x: sensorx, y: sensory}
		beacon := point{x: beaconx, y: beacony}
		grid[sensor] = "S"
		grid[beacon] = "B"
		markCoverageForSensor(sensor, beacon, grid)
	}
	for k := range grid {
		if k.x > maxx {
			maxx = k.x
		}
		if k.x < minx {
			minx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
		if k.y < miny {
			miny = k.y
		}
	}
	display(grid, minx, maxx, miny, maxy)

	notABeacon := 0
	for x := minx; x < maxx; x++ {
		if grid[point{x: x, y: row}] == "#" {
			notABeacon++
		}
	}

	fmt.Println("not a beacon: ", notABeacon)
}

func display(grid map[point]string, minx, maxx, miny, maxy int) {
	for y := miny; y <= maxy; y++ {
		fmt.Printf("%d: ", y)
		for x := minx; x <= maxx; x++ {
			if v, ok := grid[point{x: x, y: y}]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func markCoverageForSensor(sensor, beacon point, grid map[point]string) {
	maxDistance := distance(sensor.x, beacon.x, sensor.y, beacon.y)

	queue := []point{sensor}
	seen := map[point]bool{sensor: true}
	var current point

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, next := range neighbour(maxDistance, sensor, current) {
			if !seen[next] {
				seen[next] = true
				if grid[next] != "B" && grid[next] != "S" {
					grid[next] = "#"
				}
				queue = append(queue, next)
			}
		}
	}

}

var directions = []point{
	// basic directions
	{x: -1, y: 0}, // left
	{x: 0, y: -1}, // up
	{x: 1, y: 0},  // right
	{x: 0, y: 1},  // down

	// diagonally
	{x: -1, y: -1}, // left, up
	{x: 1, y: -1},  // right, up
	{x: 1, y: 1},   // right, down
	{x: -1, y: 1},  // left, down
}

func neighbour(maxDistanceFromOriginal int, original, current point) []point {
	var result []point
	for _, d := range directions {
		next := point{x: current.x + d.x, y: current.y + d.y}

		distance := distance(original.x, next.x, original.y, next.y)
		if distance <= maxDistanceFromOriginal {
			result = append(result, next)
		}
	}

	return result
}

func distance(x1, x2, y1, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
