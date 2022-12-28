package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type sensor struct {
	p        point
	coverage int
	beacon   point
}

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
	sensors := make([]sensor, 0)
	var (
		minx = math.MaxInt
		maxx int
	)
	for _, line := range split {
		var (
			sensorx, sensory int
			beaconx, beacony int
		)
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorx, &sensory, &beaconx, &beacony)
		sensorP := point{x: sensorx, y: sensory}
		beacon := point{x: beaconx, y: beacony}

		s := sensor{
			p:        sensorP,
			coverage: distance(sensorP.x, beacon.x, sensorP.y, beacon.y),
			beacon:   beacon,
		}
		sensors = append(sensors, s)

		if s.p.x-s.coverage < minx {
			minx = s.p.x - s.coverage
		}
		if s.p.x+s.coverage > maxx {
			maxx = s.p.x + s.coverage
		}
	}

	// Walk through the perimeter of each beacon's area.
	// For each sensor, construct an area of points that you check.
	// Add those points into a list of points. ( skip if below 0 or above limit )
	// For each of those points then check if they are within the
	// distance of other sensors.

	// It's a map to avoid duplicates from overlapping sensors.
	limit := 4000000
	// limit := 20
	possiblePoints := make(map[point]bool)
	for _, sensor := range sensors {
		leftCorner := point{x: sensor.p.x - sensor.coverage - 1, y: sensor.p.y}
		rightCorner := point{x: sensor.p.x + sensor.coverage + 1, y: sensor.p.y}
		upperCorner := point{x: sensor.p.x, y: sensor.p.y + sensor.coverage + 1}
		lowerCorner := point{x: sensor.p.x, y: sensor.p.y - sensor.coverage - 1}

		start := leftCorner
		possiblePoints[start] = true
		for start != lowerCorner {
			start.x++
			start.y--
			possiblePoints[start] = true
		}
		start = lowerCorner
		possiblePoints[lowerCorner] = true
		for start != rightCorner {
			start.x++
			start.y++
			possiblePoints[start] = true
		}
		start = rightCorner
		possiblePoints[start] = true
		for start != upperCorner {
			start.x--
			start.y++
			possiblePoints[start] = true
		}
		start = upperCorner
		possiblePoints[start] = true
		for start != leftCorner {
			start.x--
			start.y--
			possiblePoints[start] = true
		}
	}

	fmt.Println("length of possible points: ", len(possiblePoints))
	for k := range possiblePoints {
		if k.x < 0 || k.y < 0 || k.x > limit || k.y > limit {
			continue
		}
		inRange := false
		for _, s := range sensors {
			distanceToSensor := distance(k.x, s.p.x, k.y, s.p.y)
			if distanceToSensor <= s.coverage {
				inRange = true
				break
			}
		}
		if !inRange {
			fmt.Println("x, y: ", k)
			fmt.Println("frequency: ", k.x*limit+k.y)
			os.Exit(0)
		}
	}
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
