package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name  string
	files []File
	dirs  []string
	size  int
}

type stack struct {
	fullPath string
	s        []*Dir
}

func newStack() *stack {
	return &stack{
		s: make([]*Dir, 0),
	}
}

func (s *stack) pop() *Dir {
	var d *Dir
	d, s.s = s.s[len(s.s)-1], s.s[:len(s.s)]
	s.fullPath = filepath.Dir(s.fullPath)
	return d
}

func (s *stack) push(d *Dir) {
	s.fullPath = filepath.Join(s.fullPath, d.name)
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
		fs         = make(map[string]*Dir)
	)
	for _, v := range split {
		if v[0] == '$' {
			if strings.HasPrefix(v, "$ cd") {
				var dirName string
				fmt.Sscanf(v, "$ cd %s", &dirName)
				if dirName == ".." {
					currentDir = tree.pop()
				} else {
					if d, ok := fs[filepath.Join(tree.fullPath, dirName)]; ok {
						currentDir = d
					} else {
						currentDir = &Dir{
							name: dirName,
						}
						fs[filepath.Join(tree.fullPath, dirName)] = currentDir
					}
					tree.push(currentDir)
				}
			}
		} else {
			if strings.HasPrefix(v, "dir") {
				var dirName string
				fmt.Sscanf(v, "dir %s", &dirName)
				currentDir.dirs = append(currentDir.dirs, filepath.Join(tree.fullPath, dirName))
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
				currentDir.size += size
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
		fmt.Printf("total size of %s is: %d\n", k, size)
	}
	fmt.Println("total size: ", totalSize)
}

func calculateSize(k string, fs map[string]*Dir) int {
	size := fs[k].size
	for _, d := range fs[k].dirs {
		size += calculateSize(d, fs)
	}
	return size
}
