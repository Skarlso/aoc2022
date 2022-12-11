package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Test struct {
	// Monkey ID
	DivisibleBy int
	True        int
	False       int
}

type Monkey struct {
	ID            int
	StartingItems []int
	New           int
	Old           int
	Operation     func(a int, b int) int
	Test          Test
	Inspected     int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")

	monkey := &Monkey{}
	monkeys := make(map[int]*Monkey)
	monkeyOrder := make([]int, 0)
	for _, line := range split {
		if line == "" {
			monkeys[monkey.ID] = monkey
			monkey = &Monkey{}
			continue
		}

		if strings.HasPrefix(line, "Monkey") {
			var id int
			fmt.Sscanf(line, "Monkey %d:", &id)
			monkey.ID = id
			monkeyOrder = append(monkeyOrder, id)
		} else if strings.HasPrefix(line, "  Starting items:") {
			split := strings.Split(line, "  Starting items: ")
			ids := strings.Split(split[1], ", ")
			idsInt := make([]int, 0)
			for _, id := range ids {
				i, _ := strconv.Atoi(id)
				idsInt = append(idsInt, i)
			}
			monkey.StartingItems = idsInt
		} else if strings.HasPrefix(line, "  Operation:") {
			split := strings.Split(line, "  Operation: ")
			var (
				a, b, op   string
				aInt, bInt int
			)
			fmt.Sscanf(split[1], "new = %s %s %s", &a, &op, &b)
			if a == "old" {
				aInt = -1
			} else {
				i, _ := strconv.Atoi(a)
				aInt = i
			}
			monkey.New = aInt
			if b == "old" {
				bInt = -1
			} else {
				i, _ := strconv.Atoi(b)
				bInt = i
			}
			monkey.Old = bInt
			if op == "+" {
				monkey.Operation = func(a, b int) int {
					return a + b
				}
			} else if op == "*" {
				monkey.Operation = func(a, b int) int {
					return a * b
				}
			}
		} else if strings.HasPrefix(line, "  Test:") {
			split := strings.Split(line, "  Test: ")
			var (
				divBy int
			)
			fmt.Sscanf(split[1], "divisible by %d", &divBy)
			monkey.Test.DivisibleBy = divBy
		} else if strings.HasPrefix(line, "    If true:") {
			split := strings.Split(line, "    If true: ")
			var (
				toMonkey int
			)
			fmt.Sscanf(split[1], "throw to monkey %d", &toMonkey)
			monkey.Test.True = toMonkey
		} else if strings.HasPrefix(line, "    If false:") {
			split := strings.Split(line, "    If false: ")
			var (
				toMonkey int
			)
			fmt.Sscanf(split[1], "throw to monkey %d", &toMonkey)
			monkey.Test.False = toMonkey
		}
	}

	for i := 0; i < 20; i++ {
		for _, k := range monkeyOrder {
			fmt.Println("Looking at Monkey ", k)

			// It looks like, item is not used ATM.
			var item int
			for len(monkeys[k].StartingItems) > 0 {
				monkeys[k].Inspected++
				item, monkeys[k].StartingItems = monkeys[k].StartingItems[0], monkeys[k].StartingItems[1:]
				fmt.Println("Inspecting item: ", item)
				// increase worry level
				// somehow pass in if old is required
				var a, b int
				if monkeys[k].New > -1 {
					a = monkeys[k].New
				} else {
					a = item
				}
				if monkeys[k].Old > -1 {
					b = monkeys[k].Old
				} else {
					b = item
				}
				item = monkeys[k].Operation(a, b)
				item /= 3
				fmt.Println("Worry level is: ", item)
				if item%monkeys[k].Test.DivisibleBy == 0 {
					m := monkeys[monkeys[k].Test.True]
					m.StartingItems = append(m.StartingItems, item)
					monkeys[monkeys[k].Test.True] = m
					fmt.Printf("divisible by '%d' throwing to monkey '%d'\n", monkeys[k].Test.DivisibleBy, monkeys[k].Test.True)
				} else {
					m := monkeys[monkeys[k].Test.False]
					m.StartingItems = append(m.StartingItems, item)
					monkeys[monkeys[k].Test.False] = m
					fmt.Printf("not divisible by '%d' throwing to monkey '%d'\n", monkeys[k].Test.DivisibleBy, monkeys[k].Test.False)
				}
			}
		}
	}

	for k, v := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", k, v.Inspected)
	}

	// TODO: I manually multiplied these together. Should get it automated.
}
