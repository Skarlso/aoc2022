package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	name     string
	opString string
	op       func(a, b int) int
	eq       func(a, b int) bool
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
				m.opString = "/"
			case "+":
				m.op = func(a, b int) int { return a + b }
				m.opString = "+"
			case "-":
				m.op = func(a, b int) int { return a - b }
				m.opString = "-"
			case "*":
				m.op = func(a, b int) int { return a * b }
				m.opString = "*"
			case "=":
				m.eq = func(a, b int) bool { return a == b }
				m.opString = "="
			}
			m.a = a
			m.b = b
		} else {
			m.value = &n
		}

		monkeys[name] = m
	}

	yelled := -1
	for yelled == -1 {
		yelled = solve("root", monkeys)
		humn := monkeys["humn"]
		v := *humn.value
		v++
		humn.value = &v
		monkeys["humn"] = humn
	}
	fmt.Println("monkey yelled: ", yelled)
}

/*
comparing: a: '149' b: '150'
301 3 -
2 298 *
4 596 +
600 4 /
32 2 -
30 5 *
comparing: a: '150' b: '150'
monkey yelled:  301
*/
func solve(name string, monkeys map[string]monkey) int {
	m := monkeys[name]
	if v := m.value; v != nil {
		return *v
	}

	a := solve(m.a, monkeys)
	b := solve(m.b, monkeys)
	if m.eq != nil {
		// solve for the monkey equality function
		fmt.Printf("comparing: a: '%d' b: '%d'\n", a, b)
		if m.eq(a, b) {
			return *monkeys["humn"].value
		} else {
			return -1
		}
	}
	fmt.Println(a, b, m.opString)

	return m.op(a, b)
}
