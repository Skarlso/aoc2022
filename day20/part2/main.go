package main

import (
	"fmt"
	"os"
	"strings"
)

type number struct {
	value         int
	originalIndex int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	originalNumber := make([]int, 0)
	indexes := make(map[int]number)
	for i, line := range split {
		var n int
		fmt.Sscanf(line, "%d", &n)
		originalNumber = append(originalNumber, n)
		indexes[i] = number{value: n, originalIndex: i}
	}

	index := func(n number) int {
		for k, v := range indexes {
			if n == v {
				return k
			}
		}
		return -1
	}
	for originalIndex, n := range originalNumber {
		num := number{value: n, originalIndex: originalIndex}

		oldIndex := index(num)

		newIndex := mod((oldIndex + n), len(originalNumber)-1)

		if newIndex == 0 {
			newIndex = len(originalNumber) - 1
		}
		if newIndex >= len(originalNumber)-1 {
			newIndex = 0
		}

		oldValue := indexes[newIndex]
		indexes[newIndex] = num

		if newIndex > oldIndex {
			for k := newIndex - 1; k >= oldIndex; k-- {
				t := indexes[k]
				indexes[k] = oldValue
				oldValue = t
			}
		} else if newIndex < oldIndex {
			for k := newIndex + 1; k <= oldIndex; k++ {
				t := indexes[k]
				indexes[k] = oldValue
				oldValue = t
			}
		}
	}
	// Create a list and then wrap around after 0?
	result := make([]int, 0)
	startingIndex := 0
	for i := 0; i < len(originalNumber); i++ {
		if indexes[i].value == 0 {
			startingIndex = i
		}
		result = append(result, indexes[i].value)
	}

	fmt.Println("starting index: ", startingIndex)
	fmt.Println("result: ", result[(1000+startingIndex)%len(result)]+result[(2000+startingIndex)%len(result)]+result[(3000+startingIndex)%len(result)])
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
