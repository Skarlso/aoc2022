package main

import (
	"container/heap"
	"fmt"
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

	start := point{}
	goal := point{}

	split := strings.Split(string(content), "\n")
	for y, line := range split {
		row := make([]int, 0)
		for x, c := range line {
			if c == 'S' {
				row = append(row, 0)
				start = point{y: y, x: x}
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

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		point:    start,
		priority: 0,
	})

	cost := map[point]int{
		start: 0,
	}
	cameFrom := map[point]point{
		start: start,
	}
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		// fmt.Println(current.point, grid[current.point.y][current.point.x])
		if current.point == goal {
			break
		}
		for _, next := range neighbors(current.point, grid) {
			// the cost should be steps so far.
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

	visgrid := [][]string{}
	for y := 0; y < len(grid); y++ {
		row := make([]string, 0)
		for x := 0; x < len(grid[y]); x++ {
			row = append(row, ".")
		}
		visgrid = append(visgrid, row)
	}
	var steps int
	current := goal
	visgrid[current.y][current.x] = "E"
	for current != start {
		steps++
		current = cameFrom[current]
		visgrid[current.y][current.x] = string(rune(grid[current.y][current.x]) + 'a')
	}
	visgrid[start.y][start.x] = "S"
	fmt.Println("steps: ", steps)
	for y := 0; y < len(visgrid); y++ {
		for x := 0; x < len(visgrid[y]); x++ {
			fmt.Print(visgrid[y][x])
		}
		fmt.Println()
	}
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
