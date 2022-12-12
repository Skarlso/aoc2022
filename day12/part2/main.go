package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Item struct {
	steps    int
	point    point
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type point struct {
	x, y int
}

var (
	directions = []point{
		{x: -1, y: 0},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
	}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	grid := make([][]int, 0)

	goal := point{}
	as := make([]point, 0)

	split := strings.Split(string(content), "\n")
	for y, line := range split {
		row := make([]int, 0)
		for x, c := range line {
			if c == 'S' || c == 'a' {
				row = append(row, 0)
				as = append(as, point{y: y, x: x})
			} else if c == 'E' {
				// lowest peak I guess... let's see how that works for now.
				// maybe set it to something that it is surrounded by.
				goal = point{x: x, y: y}
				row = append(row, int('z'-'a'))
			} else {
				row = append(row, int(c-'a'))
			}
		}
		grid = append(grid, row)
	}

	minSteps := math.MaxInt
	for _, startingPoint := range as {
		pq := make(PriorityQueue, 0)
		heap.Init(&pq)
		heap.Push(&pq, &Item{
			point:    startingPoint,
			priority: 0,
		})

		cost := map[point]int{
			startingPoint: 0,
		}
		cameFrom := map[point]point{
			startingPoint: startingPoint,
		}
		found := false
		for pq.Len() > 0 {
			current := heap.Pop(&pq).(*Item)
			if current.point == goal {
				found = true
				break
			}
			for _, next := range neighbors(current.point, grid) {
				newCost := cost[current.point] + grid[next.y][next.x] + current.steps
				if v, ok := cost[next]; !ok || newCost < v {
					cameFrom[next] = current.point
					cost[next] = newCost
					heap.Push(&pq, &Item{
						point:    next,
						priority: newCost,
						steps:    current.steps + 1,
					})
				}
			}
		}

		if found {
			var steps int
			current := goal
			for current != startingPoint {
				steps++
				current = cameFrom[current]
			}

			if steps < minSteps {
				minSteps = steps
			}
		}
	}
	fmt.Println("best step: ", minSteps)
}

func neighbors(p point, grid [][]int) []point {
	var result []point
	for _, d := range directions {
		np := point{x: p.x + d.x, y: p.y + d.y}
		if np.x >= 0 && np.x < len(grid[p.y]) && np.y >= 0 && np.y < len(grid) {
			if grid[np.y][np.x] <= grid[p.y][p.x]+1 {
				result = append(result, np)
			}
		}
	}
	return result
}
