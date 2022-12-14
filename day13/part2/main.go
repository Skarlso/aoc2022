package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

	var packets []any
	for _, s := range split {
		s := strings.Split(s, "\n")
		var a, b any
		_ = json.Unmarshal([]byte(s[0]), &a)
		_ = json.Unmarshal([]byte(s[1]), &b)
		packets = append(packets, a, b)
	}

	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	decoderKey := 1
	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			decoderKey *= i + 1
		}
	}
	fmt.Println(decoderKey)
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
