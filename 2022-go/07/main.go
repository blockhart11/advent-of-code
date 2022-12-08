package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name     string
	parent   *dir
	children map[string]*dir
	files    map[string]int
	size     int
}

const (
	totalDiskSpace    = 70000000
	freeSpaceRequired = 30000000
)

func main() {
	//input, err := os.ReadFile("./2022-go/07/sample.txt")
	input, err := os.ReadFile("./2022-go/07/input.txt")
	if err != nil {
		panic(err)
	}

	var root = &dir{
		name:     "/",
		children: make(map[string]*dir),
		files:    make(map[string]int),
	}
	var wd = root

	// Let's gooooooo
	for _, v := range strings.Split(string(input), "\n") {
		cmd := strings.Split(v, " ")
		switch {
		case cmd[0] == "$" && cmd[1] == "cd":
			switch {
			case cmd[2] == "/":
				wd = root
			case cmd[2] == "..":
				if wd.parent != nil {
					wd = wd.parent
				}
			default:
				if wd.children[cmd[2]] == nil {
					wd.children[cmd[2]] = &dir{
						name:     cmd[2],
						parent:   wd,
						children: make(map[string]*dir),
						files:    make(map[string]int),
					}
				}
				wd = wd.children[cmd[2]]
			}
		case cmd[0] == "$" && cmd[1] == "ls":
			// I don't care
		case cmd[0] == "dir":
			if wd.children[cmd[1]] == nil {
				wd.children[cmd[1]] = &dir{
					name:     cmd[1],
					parent:   wd,
					children: make(map[string]*dir),
					files:    make(map[string]int),
				}
			}
		default:
			wd.files[cmd[1]], err = strconv.Atoi(cmd[0])
			if err != nil {
				panic("parsed a file wrong. doh")
			}
		}
	}

	// part 1
	computeFileSizes(root)
	fmt.Println(sumDirectoriesUnder100000(root))

	// part 2
	freeSpace := totalDiskSpace - root.size
	fmt.Printf(
		"total space: %d, required: %d, used: %d, free %d, need: %d, smallest dir: %d\n",
		totalDiskSpace,
		freeSpaceRequired,
		root.size,
		freeSpace,
		freeSpaceRequired-freeSpace,
		smallestDirSizeOverMinimum(root, freeSpaceRequired-freeSpace),
	)
}

func computeFileSizes(root *dir) {
	for _, v := range root.files {
		root.size += v
	}
	for _, v := range root.children {
		computeFileSizes(v)
		root.size += v.size
	}
}

func sumDirectoriesUnder100000(root *dir) int {
	var out int
	if root.size < 100000 {
		out += root.size
	}
	for _, v := range root.children {
		out += sumDirectoriesUnder100000(v)
	}
	return out
}

func smallestDirSizeOverMinimum(root *dir, minSize int) int {
	if root.size < minSize {
		return totalDiskSpace // should be big enough
	}
	out := root.size
	for _, v := range root.children {
		out = int(math.Min(float64(out), float64(smallestDirSizeOverMinimum(v, minSize))))
	}
	return out
}
