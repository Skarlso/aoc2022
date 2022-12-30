package main

import (
	"fmt"
	"os"
	"strings"
)

// We will be counting upwards as `+y` instead of a grid. We will handle the grid as a Cartesian coordinate system.
var (
	l = shape{
		height:   3,
		top:      point{x: 2, y: 2},
		mostLeft: point{x: 0, y: 0},
		leftSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 2, y: 1},
			{x: 2, y: 2},
		},
		rightSide: []point{
			{x: 1, y: 0},
			{x: 2, y: 0},
			{x: 2, y: 1},
			{x: 2, y: 2},
		},
		bottomSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 2, y: 0},
		},
		mostRight: point{x: 2, y: 0},
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
		height:    1,
		name:      "-",
		top:       point{x: 0, y: 0},
		mostRight: point{x: 3, y: 0},
		bottom:    point{x: 0, y: 0},
		mostLeft:  point{x: 0, y: 0},
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
		height:    2,
		name:      "B",
		top:       point{x: 0, y: 1},
		mostRight: point{x: 1, y: 0},
		bottom:    point{x: 0, y: 0},
		mostLeft:  point{x: 0, y: 0},
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
		height:    3,
		name:      "+",
		top:       point{x: 0, y: 2},
		mostRight: point{x: 1, y: 1},
		mostLeft:  point{x: -1, y: 1},
		bottom:    point{x: 0, y: 0},
		leftSide: []point{
			{x: -1, y: 1},
			{x: 0, y: 0},
			{x: 0, y: 2},
		},
		rightSide: []point{
			{x: 0, y: 0},
			{x: 1, y: 1},
			{x: 0, y: 2},
		},
		bottomSide: []point{
			{x: -1, y: 1},
			{x: 0, y: 0},
			{x: 1, y: 1},
		},
		all: []point{
			{x: 0, y: 0},
			{x: -1, y: 1},
			{x: 0, y: 2},
			{x: 1, y: 1},
			{x: 0, y: 1},
		},
	}
	linePiece = shape{
		height:    4,
		name:      "|",
		top:       point{x: 0, y: 3},
		mostRight: point{x: 0, y: 0},
		bottom:    point{x: 0, y: 0},
		mostLeft:  point{x: 0, y: 0},
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
	height     int
	leftSide   []point
	rightSide  []point
	top        point
	mostRight  point
	mostLeft   point
	bottom     point
	bottomSide []point
	all        []point

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

	y := 0
	rocks := 2022
	fallen := 0
	playground := make(map[point]bool)
	leftSide := 0
	rightSide := 7
	jetPush := 0
	for fallen != rocks {
		currentRock := fallingOrder[fallen%len(fallingOrder)]
		startingY := y + 3
		current := point{x: 2 - currentRock.mostLeft.x, y: startingY} // This is the location of the designated zeroth coordinate of the shape.
		canMove := true
		for canMove {
			canMove = false

			currentJet := jetPattern[jetPush%len(jetPattern)]
			if currentJet == "<" {
				if current.x-1+currentRock.mostLeft.x >= leftSide && !isSomethingToTheLeft(current, currentRock, playground) {
					current.x--
				}
			} else if currentJet == ">" {
				if current.x+1+currentRock.mostRight.x < rightSide && !isSomethingToTheRight(current, currentRock, playground) {
					current.x++
				}
			}
			jetPush++

			if (current.y-1+currentRock.bottom.y) >= 0 && !isSomethingDownwards(current, currentRock, playground) {
				current.y--
				canMove = true
			}
		}

		for _, p := range allPoints(current, currentRock) {
			playground[p] = true
		}
		if current.y+currentRock.height > y {
			y = current.y + currentRock.height
		}

		fallen++
	}

	fmt.Println("The tower is this tall: ", y)
}

func display(playground map[point]bool, maxy int, current point, rock shape) {
	contains := func(p point, current point, list []point) bool {
		for _, v := range list {
			if p.x == v.x+current.x && p.y == v.y+current.y {
				return true
			}
		}

		return false
	}
	for y := maxy + 7; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if playground[point{x: x, y: y}] {
				fmt.Print("#")
			} else if contains(point{x: x, y: y}, current, rock.all) {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("-------")
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
	for _, p := range currentRock.bottomSide {
		if playground[point{x: p.x + current.x, y: current.y + p.y - 1}] {
			return true
		}
	}
	return false
}
