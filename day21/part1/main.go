package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	name string
	op   func(a, b int) int
	// if not nil, that means we need to look at the operation to compute it
	value *int
	// there are no situations in which it's a name and a number. It's always name OP name
	a string
	b string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	monkeys := make(map[string]monkey)
	for _, line := range split {
		split := strings.Split(line, ": ")
		name := split[0]
		m := monkey{
			name: name,
		}
		n, err := strconv.Atoi(split[1])
		if err != nil {
			var (
				a, op, b string
			)
			fmt.Sscanf(split[1], "%s %s %s", &a, &op, &b)
			// fmt.Printf("a: '%s', b: '%s', op: '%s'\n", a, b, op)
			switch op {
			case "/":
				m.op = func(a, b int) int { return a / b }
			case "+":
				m.op = func(a, b int) int { return a + b }
			case "-":
				m.op = func(a, b int) int { return a - b }
			case "*":
				m.op = func(a, b int) int { return a * b }
			}
			m.a = a
			m.b = b
		} else {
			m.value = &n
		}

		monkeys[name] = m
	}

	yelled := solve("root", monkeys)
	fmt.Println("monkey yelled: ", yelled)
}

func solve(name string, monkeys map[string]monkey) int {
	m := monkeys[name]
	if v := m.value; v != nil {
		return *v
	}

	a := solve(m.a, monkeys)
	b := solve(m.b, monkeys)

	return m.op(a, b)
}
