package main

import (
	"fmt"
	"math"
	"os"
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
	n := 0
	digits := map[string]int{
		"2": 2, "1": 1, "0": 0, "-": -1, "=": -2,
	}
	numbers := map[int]string{
		2: "2", 1: "1", 0: "0", 3: "=", 4: "-", 5: "0",
	}

	for _, line := range split {
		chars := strings.Split(line, "")
		reverse(chars)
		sum := 0
		for i, digit := range chars {
			sum += int(math.Pow(5, float64(i))) * digits[digit]
		}
		n += sum
	}
	snafu, carry := []string{}, 0
	for n > 0 {
		x := (n % 5) + carry
		snafu = append(snafu, numbers[x])
		if x > 2 {
			carry = 1
		} else {
			carry = 0
		}
		n /= 5
	}

	reverse(snafu)
	fmt.Println(strings.Join(snafu, ""))
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
