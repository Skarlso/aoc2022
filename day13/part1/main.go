package main

import (
	"encoding/json"
	"fmt"
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
	split := strings.Split(strings.TrimSpace(string(content)), "\n\n")

	var ordered int
	for i, s := range split {
		s := strings.Split(s, "\n")
		var a, b any
		_ = json.Unmarshal([]byte(s[0]), &a)
		_ = json.Unmarshal([]byte(s[1]), &b)
		if cmp(a, b) <= 0 {
			ordered += i + 1
		}
	}
	fmt.Println(ordered)
}

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}
