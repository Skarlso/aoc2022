package main

import (
	"fmt"
	"os"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name  string
	files []File
	Dirs  []identifier
	size  int
}

type identifier struct {
	depth int
	name  string
}

type stack struct {
	depth int
	s     []*Dir
}

func newStack() *stack {
	return &stack{
		s: make([]*Dir, 0),
	}
}

func (s *stack) pop() *Dir {
	var d *Dir
	d, s.s = s.s[len(s.s)-1], s.s[:len(s.s)]
	s.depth++
	return d
}

func (s *stack) push(d *Dir) {
	s.depth--
	s.s = append(s.s, d)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")

	tree := newStack()
	var (
		currentDir *Dir
		fs         = make(map[identifier]*Dir)
	)
	for _, v := range split {
		if v[0] == '$' {
			if strings.HasPrefix(v, "$ cd") {
				var dirName string
				fmt.Sscanf(v, "$ cd %s", &dirName)
				if dirName == ".." {
					currentDir = tree.pop()
				} else {
					if d, ok := fs[identifier{
						name:  dirName,
						depth: tree.depth,
					}]; ok {
						currentDir = d
					} else {
						currentDir = &Dir{
							name: dirName,
						}
						fs[identifier{
							name:  dirName,
							depth: tree.depth,
						}] = currentDir
					}
					tree.push(currentDir)
				}
			}
		} else {
			if strings.HasPrefix(v, "dir") {
				var dirName string
				fmt.Sscanf(v, "dir %s", &dirName)
				currentDir.Dirs = append(currentDir.Dirs, identifier{name: dirName, depth: tree.depth})
			} else {
				var (
					name string
					size int
				)
				fmt.Sscanf(v, "%d %s", &size, &name)
				currentDir.files = append(currentDir.files, File{
					name: name,
					size: size,
				})
				// currentDir.size += size
			}
		}
	}

	limit := 100000
	totalSize := 0
	for k := range fs {
		size := calculateSize(k, fs)
		if size <= limit {
			totalSize += size
		}
		fmt.Printf("total size of %s is: %d\n", k.name, size)
	}
	fmt.Println("total size: ", totalSize)
}

func calculateSize(k identifier, fs map[identifier]*Dir) int {
	size := 0
	for _, f := range fs[k].files {
		size += f.size
	}
	for _, d := range fs[k].Dirs {
		size += calculateSize(d, fs)
	}
	return size
}
