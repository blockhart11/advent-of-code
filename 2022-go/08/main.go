package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	height int
	// highest tree height to edge
	up, down, left, right int
	// visible in any direction?
	any bool
	// number of trees visible in all directions
	numUp, numDown, numLeft, numRight int
}

func main() {
	//input, err := os.ReadFile("./2022-go/08/sample.txt")
	input, err := os.ReadFile("./2022-go/08/input.txt")
	if err != nil {
		panic(err)
	}

	// create the forest
	rows := strings.Split(string(input), "\n")
	forest := make([][]*tree, len(rows))
	for i, row := range strings.Split(string(input), "\n") {
		forest[i] = make([]*tree, len(row))
		for j, col := range row {
			h, err := strconv.Atoi(string(col))
			if err != nil {
				panic("oh noes!")
			}
			forest[i][j] = &tree{height: h}
		}
	}

	// First go left to right, top to bottom
	for i, row := range forest {
		for j, _ := range row {
			// check up and left
			checkVisibleLeft(forest, i, j)
			checkVisibleUp(forest, i, j)
			// Part 2 - might as well do it here...
			countVisibleLeft(forest, i, j)
			countVisibleUp(forest, i, j)
			countVisibleRight(forest, i, j)
			countVisibleDown(forest, i, j)
		}
	}
	// now go right to left, bottom to top
	for i := len(forest) - 1; i >= 0; i-- {
		for j := len(forest[i]) - 1; j >= 0; j-- {
			checkVisibleRight(forest, i, j)
			checkVisibleDown(forest, i, j)
		}
	}
	// now count them
	var countVisible int
	for _, row := range forest {
		for _, col := range row {
			if col.any {
				fmt.Printf("1")
				countVisible++
			} else {
				fmt.Printf("0")
			}
		}
		fmt.Printf("\n")
	}
	// Part 1
	fmt.Println(countVisible)

	// Part 2
	var maxScenicScore int
	for i, row := range forest {
		for j, _ := range row {
			t := forest[i][j]
			score := t.numLeft * t.numUp * t.numRight * t.numDown
			if maxScenicScore < score {
				maxScenicScore = score
				fmt.Printf("(%d, %d): %d -- left: %d, up: %d, right: %d, down: %d\n", i, j, score, t.numLeft, t.numUp, t.numRight, t.numDown)
				fmt.Println("New high score!")
			}
		}
	}
	fmt.Println(maxScenicScore)
}

func checkVisibleLeft(forest [][]*tree, row, col int) {
	switch {
	case col == 0:
		forest[row][col].left = forest[row][col].height
		forest[row][col].any = true
	case forest[row][col].height > forest[row][col-1].left:
		forest[row][col].left = forest[row][col].height
		forest[row][col].any = true
	default:
		forest[row][col].left = forest[row][col-1].left
	}
}

func checkVisibleUp(forest [][]*tree, row, col int) {
	switch {
	case row == 0:
		forest[row][col].up = forest[row][col].height
		forest[row][col].any = true
	case forest[row][col].height > forest[row-1][col].up:
		forest[row][col].up = forest[row][col].height
		forest[row][col].any = true
	default:
		forest[row][col].up = forest[row-1][col].up
	}
}

func checkVisibleRight(forest [][]*tree, row, col int) {
	switch {
	case col == len(forest[row])-1:
		forest[row][col].right = forest[row][col].height
		forest[row][col].any = true
	case forest[row][col].height > forest[row][col+1].right:
		forest[row][col].right = forest[row][col].height
		forest[row][col].any = true
	default:
		forest[row][col].right = forest[row][col+1].right
	}
}

func checkVisibleDown(forest [][]*tree, row, col int) {
	switch {
	case row == len(forest)-1:
		forest[row][col].down = forest[row][col].height
		forest[row][col].any = true
	case forest[row][col].height > forest[row+1][col].down:
		forest[row][col].down = forest[row][col].height
		forest[row][col].any = true
	default:
		forest[row][col].down = forest[row+1][col].down
	}
}

func countVisibleLeft(forest [][]*tree, row, col int) {
	if col == 0 {
		return
	}
	var count int
	for i := col - 1; i >= 0; i-- {
		count++
		if forest[row][i].height >= forest[row][col].height {
			break
		}
	}
	forest[row][col].numLeft = count
}

func countVisibleUp(forest [][]*tree, row, col int) {
	if row == 0 {
		return
	}
	var count int
	for i := row - 1; i >= 0; i-- {
		count++
		if forest[i][col].height >= forest[row][col].height {
			break
		}
	}
	forest[row][col].numUp = count
}

func countVisibleRight(forest [][]*tree, row, col int) {
	if col == len(forest[row])-1 {
		return
	}
	var count int
	for i := col + 1; i < len(forest[row]); i++ {
		count++
		if forest[row][i].height >= forest[row][col].height {
			break
		}
	}
	forest[row][col].numRight = count
}

func countVisibleDown(forest [][]*tree, row, col int) {
	if row == len(forest)-1 {
		return
	}
	var count int
	for i := row + 1; i < len(forest); i++ {
		count++
		if forest[i][col].height >= forest[row][col].height {
			break
		}
	}
	forest[row][col].numDown = count
}
