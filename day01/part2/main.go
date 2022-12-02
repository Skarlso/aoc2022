package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	max1 := 0
	max2 := 0
	max3 := 0
	sumSoFar := 0
	for _, l := range split {
		if l == "" {
			if sumSoFar > max1 {
				if max2 > max3 {
					max3 = max2
				}
				if max1 > max2 {
					max2 = max1
				}
				if max1 < max2 && max1 > max3 {
					max3 = max1
				}
				max1 = sumSoFar
			}
			if sumSoFar < max1 && sumSoFar > max2 {
				if max2 > max3 {
					max3 = max2
				}
				max2 = sumSoFar
			}
			if sumSoFar < max2 && sumSoFar > max3 {
				max3 = sumSoFar
			}
			sumSoFar = 0
			continue
		}
		i, _ := strconv.Atoi(strings.Trim(l, "\n"))
		sumSoFar += i
	}
	fmt.Println(max1 + max2 + max3)
}
