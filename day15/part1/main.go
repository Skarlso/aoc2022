package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

	notABeacon := 0
	for x := minx; x <= maxx; x++ {
		p := point{x: x, y: row}
		inRange := false
		for _, s := range sensors {
			distanceToSensor := distance(p.x, s.p.x, p.y, s.p.y)
			if distanceToSensor <= s.coverage {
				inRange = true
				break
			}
		}

		if inRange {
			notABeacon++
		}
	}

	// -1 for S
	fmt.Println("not a beacon: ", notABeacon-1)
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
