package main

import (
	"fmt"
	"os"
	"strings"
)

type cpu struct {
	register     int
	clockCircuit int
	width        int
}

func (c *cpu) addx(v int) {
	for i := 0; i < 2; i++ {
		if c.clockCircuit == c.register-1 || c.clockCircuit == c.register || c.clockCircuit == c.register+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		c.clockCircuit++
		if c.clockCircuit == c.width {
			c.clockCircuit = 0
			fmt.Println()
		}
	}
	c.register += v
}

func (c *cpu) noop() {
	if c.clockCircuit == c.register-1 || c.clockCircuit == c.register || c.clockCircuit == c.register+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	c.clockCircuit++
	if c.clockCircuit == c.width {
		c.clockCircuit = 0
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	c := cpu{register: 1, width: 40}

	for _, line := range split {
		var (
			op    string
			value int
		)
		fmt.Sscanf(line, "%s %d", &op, &value)

		switch op {
		case "addx":
			c.addx(value)
		case "noop":
			c.noop()
		}
	}
}
