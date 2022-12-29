package main

import (
	"fmt"
	"os"
	"strings"
)

// We will be counting upwards as `+y` instead of a grid. We will handle the grid as a Cartesian coordinate system.
var (
	l = shape{
		top: point{x: 2, y: 2},
		leftSide: []point{
			{x: 0, y: 0},
			{x: 2, y: 1},
			{x: 2, y: 2},
		},
		rightSide: []point{
			{x: 2, y: 0},
			{x: 2, y: 1},
			{x: 2, y: 2},
		},
		bottomSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 2, y: 0},
		},
		mostRight: point{x: 1, y: 0},
		bottom:    point{x: 0, y: 0},
		all: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 2, y: 0},
			{x: 2, y: 1},
			{x: 2, y: 2},
		},
		name: "L",
	}
	minus = shape{
		name:      "-",
		top:       point{x: 0, y: 0},
		mostRight: point{x: 3, y: 0},
		bottom:    point{x: 0, y: 0},
		leftSide: []point{
			{x: 0, y: 0},
		},
		rightSide: []point{
			{x: 3, y: 0},
		},
		bottomSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 2, y: 0},
			{x: 3, y: 0},
		},
		all: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 2, y: 0},
			{x: 3, y: 0},
		},
	}
	block = shape{
		name:      "B",
		top:       point{x: 0, y: 1},
		mostRight: point{x: 1, y: 0},
		bottom:    point{x: 0, y: 0},
		leftSide: []point{
			{x: 0, y: 0},
			{x: 0, y: 1},
		},
		rightSide: []point{
			{x: 1, y: 0},
			{x: 1, y: 1},
		},
		bottomSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
		},
		all: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 0, y: 1},
			{x: 1, y: 1},
		},
	}
	plus = shape{
		name:      "+",
		top:       point{x: 1, y: 1},
		mostRight: point{x: 2, y: 0},
		bottom:    point{x: 1, y: -1},
		leftSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 1},
			{x: 1, y: -1},
		},
		rightSide: []point{
			{x: 1, y: 1},
			{x: 1, y: -1},
			{x: 2, y: 0},
		},
		bottomSide: []point{
			{x: 0, y: 0},
			{x: 1, y: -1},
			{x: 2, y: 0},
		},
		all: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 1, y: -1},
			{x: 1, y: 1},
			{x: 2, y: 0},
		},
	}
	linePiece = shape{
		name:      "|",
		top:       point{x: 0, y: 3},
		mostRight: point{x: 0, y: 0},
		bottom:    point{x: 0, y: 0},
		leftSide: []point{
			{x: 0, y: 0},
			{x: 0, y: 1},
			{x: 0, y: 2},
			{x: 0, y: 3},
		},
		rightSide: []point{
			{x: 0, y: 0},
			{x: 0, y: 1},
			{x: 0, y: 2},
			{x: 0, y: 3},
		},
		bottomSide: []point{
			{x: 0, y: 0},
		},
		all: []point{
			{x: 0, y: 0},
			{x: 0, y: 1},
			{x: 0, y: 2},
			{x: 0, y: 3},
		},
	}
)

// Share is made up of a set of points that we will add to the current grid.
// They will be used to continuously calculate where each point of a shape is
// located at at any given moment in the movement.
// Falling order:
// ####
//
// .#.
// ###
// .#.
//
// ..#
// ..#
// ###
//
// #
// #
// #
// #
//
// ##
// ##
type shape struct {
	leftSide   []point
	rightSide  []point
	top        point
	mostRight  point
	bottom     point
	bottomSide []point
	all        []point
	// points []point
	// bottom
	// top

	name string
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

	fallingOrder := []shape{minus, plus, l, linePiece, block}
	jetPattern := make([]string, 0)
	for _, line := range split {
		jetPattern = strings.Split(line, "")
	}
	// fmt.Println(fallingOrder)
	fmt.Println("jet pattern: ", jetPattern)

	y := 4
	unitsOfTall := 0
	// endingPoint := 0

	rocks := 2
	fallen := 0
	playground := make(map[point]bool)
	leftSide := 0
	rightSide := 7
	jetPush := 0
	display(playground, y)
	for fallen != rocks {
		currentRock := fallingOrder[fallen%len(fallingOrder)]
		// fmt.Println("current falling rock: ", currentRock.name)
		// make sure we are 3 above the last fallen rock's last y coordinate considering our current falling rock's
		// lowest point.
		y += currentRock.bottom.y
		start := point{x: 2, y: y} // This is the location of the designated zeroth coordinate of the shape.
		// lowestPoint := lowest(rockCoordinate, currentRock)
		// // Remember that this is +y because we are in a Cartesian coordinate system, I decided.
		// highestPoint := highest(rockCoordinate, currentRock)

		queue := []point{start}
		var current point
		for len(queue) > 0 {
			fmt.Println("moving")
			display(playground, y)
			// Determine the coordinate of each piece of the shape as it falls down.
			// let it fall until any of its coordinates hits anything
			current, queue = queue[0], queue[1:]

			// move down
			// if we didn't reach the end and we didn't hit anything on the way down,
			// we increase the x, y coordinate which will track our rock across the movement.
			if (current.y-1+currentRock.bottom.y) > 0 && !isSomethingDownwards(current, currentRock, playground) {
				// add it to the new round, and then push the block.
				queue = append(queue, point{x: current.x, y: current.y - 1})
				current.y--
				// if we hit anything, the push still happens, but it might just not move it. however, the push
				// might have pushed it into a direction that now lets it fall.
				// so once we push we check again if anything is below us. If not, we include this check again.
			}

			// push the block

			currentJet := jetPattern[jetPush%len(jetPattern)]
			if currentJet == "<" {
				// The x is at the left side, so this is okay.
				// But I must also consider all points because it could blow the rock into a crevasse.
				if current.x-1 >= leftSide && !isSomethingToTheLeft(current, currentRock, playground) {
					current.x--
				}
			} else if currentJet == ">" {
				if current.x+currentRock.mostRight.x < rightSide && !isSomethingToTheRight(current, currentRock, playground) {
					current.x++
				}
			}
			jetPush++

			// check again if something is below us, if not, we add it to the queue, but we don't decrease y again
			if (current.y-1+currentRock.bottom.y) > 0 && !isSomethingDownwards(current, currentRock, playground) {
				// add it to the new round, and then push the block.
				queue = append(queue, point{x: current.x, y: current.y})
			}
		}

		// Once the rock stopped moving we add each point of it to the playground.
		for _, p := range allPoints(current, currentRock) {
			playground[p] = true
		}

		// shit, I need to increase TALL
		// Maybe this is fine.
		// increase the y at which we start.
		if current.y+currentRock.top.y > y {
			y += current.y + currentRock.top.y
			// I have to re-think this one.
			unitsOfTall += (current.y + currentRock.top.y - unitsOfTall)
		}

		fallen++
	}

	fmt.Println("The tower is this tall: ", unitsOfTall)
}

func display(playground map[point]bool, maxy int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < 7; x++ {
			if playground[point{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func allPoints(current point, rock shape) []point {
	var calculated []point
	for _, p := range rock.all {
		calculated = append(calculated, point{x: current.x + p.x, y: current.y + p.y})
	}

	return calculated
}

func isSomethingToTheLeft(current point, currentRock shape, playground map[point]bool) bool {
	for _, p := range currentRock.leftSide {
		// -1 because we are checking the neighbour
		if playground[point{x: p.x + current.x - 1, y: current.y + p.y}] {
			return true
		}
	}
	return false
}

func isSomethingToTheRight(current point, currentRock shape, playground map[point]bool) bool {
	// +1 because we are checking the neighbour
	for _, p := range currentRock.rightSide {
		if playground[point{x: p.x + current.x + 1, y: current.y + p.y}] {
			return true
		}
	}
	return false
}

// include the bottom line which is y == 0.
func isSomethingDownwards(current point, currentRock shape, playground map[point]bool) bool {
	// -1 because we are checking the neighbour
	for _, p := range currentRock.rightSide {
		if playground[point{x: p.x + current.x, y: current.y + p.y - 1}] {
			return true
		}
	}
	return false
}
