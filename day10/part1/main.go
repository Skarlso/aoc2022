package main

import (
	"fmt"
	"os"
	"strings"
)

type cpu struct {
	register     int
	clockCircuit int
	// mark tracks the 40th tick
	mark int
	// signal strength
	signalStrength int
}

func (c *cpu) addx(v int) {
	for i := 0; i < 2; i++ {
		// check cycle count.
		c.clockCircuit++
		if c.mark == c.clockCircuit {
			// fmt.Printf("mark %d. register: %d\n", c.mark, c.register)
			c.signalStrength += c.mark * c.register
			c.mark = c.clockCircuit + 40
			// fmt.Println("signal strength: ", c.signalStrength)
		}
	}
	c.register += v
}

func (c *cpu) noop() {
	c.clockCircuit++
	if c.mark == c.clockCircuit {
		// fmt.Printf("mark %d. register: %d\n", c.mark, c.register)
		c.signalStrength += c.mark * c.register
		c.mark = c.clockCircuit + 40
		// fmt.Println("signal strength: ", c.signalStrength)
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
	c := cpu{register: 1, mark: 20}

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

	fmt.Println("signal strength: ", c.signalStrength)
}
